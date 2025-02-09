package main

//the comparison takes advantage of Nick Lebesis' gabbra-train-v1 from Huggingface
import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo" //for discord communication
	"github.com/joho/godotenv"      //for env files (API KEY)
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type InputMessage struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	User      string `json:"user"`
}

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

		writeFile(m.Content, m.Timestamp.String(), m.Author.String())

		switch {
		case strings.Contains(strings.ToLower(m.Content), "hi"): //change out for commands
			discord.ChannelMessageSend(m.ChannelID, "Heyyyyyy")
		case strings.Contains(strings.ToLower(m.Content), "Grogu"):
			discord.ChannelMessageSend(m.ChannelID, "Grogu is my name!!")
		case strings.Contains(strings.ToLower(m.Content), "/analyze_insights"): //command to analyze community insights
			analyseInsights()
		case strings.Contains(strings.ToLower(m.Content), "/help"):
			discord.ChannelMessageSend(m.ChannelID, "Commands:\n/analyse_insights returns server insights")
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

// analyseInsights analyzes the discord community's insights and presents them to the user
func analyseInsights() {

	//open file
	dataset, err := os.ReadFile("dataset.json")
	prompt := ("Please provide a summary of important people and their key topics in this collection of discord messages: " + (json.Unmarshal(dataset, &message).String()))
	python := path.Clean(strings.Join([]string{os.Getenv("userprofile"), "Anaconda3", "python.exe"}, "/"))
	script := "torus.py"
	cmd := exec.Command("cmd", python, script, prompt)
	out, err := cmd.Output()
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}
	//compare it

	//print insights

}

func writeFile(content string, timestamp string, author string) {
	input := InputMessage{
		Message:   content,
		Timestamp: timestamp,
		User:      author,
	}

	// Initialize an empty slice to hold the input messages
	var inputs []InputMessage

	// Try to read the existing data from the file
	fileData, err := os.ReadFile("data.json")
	if err == nil && len(fileData) > 0 {
		// If file exists and contains data, unmarshal it into inputs slice
		err = json.Unmarshal(fileData, &inputs)
		if err != nil {
			fmt.Println("Error unmarshaling data:", err)
			return
		}
	}

	// Append the new message to the slice
	inputs = append(inputs, input)

	// Convert the updated inputs slice back to JSON
	jsonData, err := json.MarshalIndent(inputs, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// Save the updated data back to the file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Inputs saved successfully")
}
