package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// HuggingFaceResponse represents the response from Hugging Face
type HuggingFaceResponse struct {
	GeneratedText string `json:"generated_text"`
}

// AnalyzeMessage sends a message to Hugging Face and gets a response
func AnalyzeMessage(message string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiURL := "https://api-inference.huggingface.co/models/facebook/blenderbot-400M-distill"
	apiToken := os.Getenv("HUGGINGFACE_API_TOKEN")

	payload := map[string]string{"inputs": message}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "I'm having trouble understanding right now. Try again later."
	}
	defer resp.Body.Close()

	var response []HuggingFaceResponse
	json.NewDecoder(resp.Body).Decode(&response)
	if len(response) > 0 {
		return response[0].GeneratedText
	}
	return "I'm not sure how to respond to that."
}
