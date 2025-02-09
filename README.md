Overview
In this project, we developed a Discord bot in Go that reads and stores users' messages, along with their associated timestamp and user information. The bot was designed to monitor Discord server activity and collect message data for future analysis.

While we began the process of analyzing the stored messages using Torus, a network designed to securely store and analyze data, the analysis was not fully completed. The plan was to use Torus to generate summaries and statistics about the messages, as well as identify key users in the conversation.

Additionally, we aimed to compare the stored messages with a publicly available dataset, Nick Lebesis' gabbra-train-v1 from Huggingface, in order to flag inappropriate behavior. By evaluating user messages in context, we planned to identify and flag any malicious or predatory behavior based on comparison with the dataset.

Features
Message Logging: The bot logs each user message along with the timestamp and author.
Torus Integration (Partially Implemented): The bot was designed to use the Torus network to securely store user messages and generate message summaries and statistics.
Malicious Behavior Detection (Planned): The bot aims to compare messages with a dataset (Nick Lebesis' gabbra-train-v1) to flag inappropriate or malicious behavior based on context.
Technologies Used
Go (Golang): The primary programming language used to develop the bot.
DiscordGo: A Go library for interacting with the Discord API.
Torus: A network used to securely store data and analyze it.
Huggingface Dataset: Nick Lebesis' gabbra-train-v1 dataset used to flag inappropriate behavior.
