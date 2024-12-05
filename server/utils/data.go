package utils

import (
	"encoding/json"
	"os"
	"strings"
)

// FAQ represents a frequently asked question and its answer
type FAQ struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// LoadFAQs loads FAQ data from a JSON file
func LoadFAQs() ([]FAQ, error) {
	file, err := os.Open("data/healthcare_faqs.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var faqs []FAQ
	err = json.NewDecoder(file).Decode(&faqs)
	return faqs, err
}

// FindFAQ searches for a matching FAQ answer
func FindFAQ(message string) string {
	faqs, err := LoadFAQs()
	if err != nil {
		return "Sorry, I couldn't load FAQs right now. Try again later."
	}

	for _, faq := range faqs {
		if strings.Contains(strings.ToLower(message), strings.ToLower(faq.Question)) {
			return faq.Answer
		}
	}
	return ""
}
