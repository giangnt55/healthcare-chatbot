package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// OpenAIRequest represents the payload for ChatGPT
type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// ChatMessage represents a single message in the ChatGPT conversation
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response from ChatGPT
type OpenAIResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

// AnalyzeWithChatGPT sends a request to ChatGPT API and retrieves the response
func AnalyzeWithChatGPT(question string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiURL := "https://api.openai.com/v1/chat/completions"
	apiToken := os.Getenv("OPENAI_API_KEY")

	// Prepare ChatGPT payload
	payload := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{Role: "system", Content: "You are a helpful healthcare assistant."},
			{Role: "user", Content: question},
		},
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return "Error connecting to ChatGPT API."
	}
	defer resp.Body.Close()

	// Handle non-OK responses
	if resp.StatusCode != http.StatusOK {
		var errResponse map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errResponse)
		log.Printf("Error response from ChatGPT API: %v", errResponse)
		return fmt.Sprintf("ChatGPT API error: %v", errResponse["error"])
	}

	// Parse ChatGPT response
	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error decoding ChatGPT response: %v", err)
		return "Error decoding the ChatGPT response."
	}

	// Extract the first response choice
	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content
	}

	return "No response from ChatGPT."
}
