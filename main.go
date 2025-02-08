package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Create a new Discord session
	discord, err := discordgo.New("Bot " + os.Getenv("API_KEY"))
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Add event handlers
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready!")
	})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Ignore messages from the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		switch {
		case strings.Contains(strings.ToLower(m.Content), "hi"): //change out for commands
			discord.ChannelMessageSend(m.ChannelID, "Heyyyyyy")
		case strings.Contains(strings.ToLower(m.Content), "Grogu"):
			discord.ChannelMessageSend(m.ChannelID, "Grogu is my name!!")
		case strings.Contains(strings.ToLower(m.Content), "analyze insights"):

		}

		// Print received message
		fmt.Println("Message received:", m.Content)

		// Send a reply message in the chat
		s.ChannelMessageSend(m.ChannelID, "Hello! I received your message: "+m.Content)
	})

	// Open the connection
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	// Keep the bot running
	select {}
}
