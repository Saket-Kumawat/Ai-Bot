package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Edw590/go-wolfram"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient = &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}

	go printCommandEvents(bot.CommandEvents())

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{
		Description: "Send any question to Wolfram Alpha",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")

			// Step 1: Send to Wit.ai
			msg, err := client.Parse(&witai.MessageRequest{Query: query})
			if err != nil {
				log.Printf("Wit.ai error: %v", err)
				response.Reply("Error processing your message with Wit.ai.")
				return
			}

			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])

			// Step 2: Try to extract Wolfram query entity
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")

			var answer string
			if value.Exists() {
				answer = value.String()
			} else {
				log.Println("No valid entity found, falling back to user input.")
				answer = query // Fallback to raw input
			}

			fmt.Printf("Query to Wolfram Alpha: %s\n", answer)

			// Step 3: Ask Wolfram
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				log.Printf("Wolfram error: %v", err)
				response.Reply("Wolfram Alpha did not understand your input.")
				return
			}

			response.Reply(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil {
		log.Fatalf("Slack bot failed to start: %v", err)
	}
}
