package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"healthcare-chatbot/models"
	"healthcare-chatbot/utils"
)

// ChatHandler processes user messages
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var userMessage struct {
		Message string `json:"message"`
	}
	json.NewDecoder(r.Body).Decode(&userMessage)

	// First, try to match an FAQ response
	response := utils.FindFAQ(userMessage.Message)
	if response == "" {
		// If no FAQ matches, use Hugging Face NLP
		response = utils.AnalyzeMessage(userMessage.Message)
	}
	fmt.Sprintf("Message: %s - Response: %s", userMessage.Message, response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ChatResponse{Message: response})
}
