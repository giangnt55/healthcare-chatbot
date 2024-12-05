package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// FAQ represents a frequently asked question and its answer
type FAQ struct {
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
}

// LoadFAQs loads FAQ data from a JSON file
func LoadFAQs() ([]FAQ, error) {
	var queries []FAQ
	file, err := os.Open("data/healthcare_faqs.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse each line as a JSON object into MedicalQuery
		var query FAQ
		err := json.Unmarshal([]byte(scanner.Text()), &query)
		if err != nil {
			return nil, fmt.Errorf("error parsing JSON: %v", err)
		}
		// Append the parsed query to the slice
		queries = append(queries, query)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}

	return queries, nil
}

// FindFAQ searches for a matching FAQ answer
func FindFAQ(message string) string {
	return ""
	queries, err := LoadFAQs()
	if err != nil {
		return "Sorry, I couldn't load medical queries right now. Try again later."
	}

	for _, query := range queries {
		// We will search for a match based on the input description
		if strings.Contains(strings.ToLower(query.Input), strings.ToLower(message)) {
			return query.Output
		}
	}

	// If no match is found, return a default response
	return ""
}
