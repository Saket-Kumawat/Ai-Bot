🤖 Slack AI Bot with Wit.ai & Wolfram Alpha
This project is a Slack-integrated AI chatbot written in Go that leverages Wit.ai for Natural Language Understanding (NLU) and Wolfram Alpha for knowledge-based responses. It listens to specific messages in Slack and returns intelligent answers in real time.

📌 Features
🧠 Wit.ai Integration: Extracts meaning and intent from user messages.

🔍 Wolfram Alpha API: Provides accurate, knowledge-rich answers.

💬 Slack Real-Time Messaging: Responds directly in Slack using Socket Mode.

📦 Built using the lightweight and fast Gin + Slacker frameworks.

💡 Fallback handling when NLP fails (uses raw user query directly).

Theoretical Overview
💬 Natural Language Processing with Wit.ai
Wit.ai is used to process the text input from the user and extract structured data such as entities and intents.

It converts unstructured human language into structured JSON.

Trained entities (like wit$wolfram_search_query) allow the bot to understand specific queries.

If Wit.ai cannot detect the proper intent/entity, the raw user input is used as a fallback.

Why NLP matters: Traditional bots work with fixed commands; NLP allows understanding flexible, natural sentences like:

“Can you tell me the distance from Earth to Mars?”

🔍 Knowledge Retrieval via Wolfram Alpha API
Wolfram Alpha is a computational knowledge engine capable of answering factual, mathematical, and scientific questions.

The bot forwards the query to the Wolfram API using the GetSpokentAnswerQuery() method.

Wolfram returns a spoken-style answer, ready for display to the user.

This turns the bot into a factual answering machine for general knowledge.

Example queries it can answer:

What is the square root of 144?

Capital of Japan?

Weather in New York?

Speed of light in km/s?

