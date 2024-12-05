package utils

import (
	"fmt"
	"io/ioutil"
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

	apiURL := "https://huggingface.co/datasets/wangrongsheng/HealthCareMagic-100k-en/1.0.0/en"
	apiToken := os.Getenv("HUGGINGFACE_API_TOKEN")

	// Tạo yêu cầu HTTP GET
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Thêm header Authorization với token API của bạn
	req.Header.Add("Authorization", "Bearer "+apiToken)

	// Tạo client và gửi yêu cầu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Đọc và xử lý dữ liệu trả về
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// In dữ liệu ra màn hình
	fmt.Println("Response body:", string(body))

	// payload := map[string]string{"inputs": message}
	// body, _ := json.Marshal(payload)

	// req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	// req.Header.Set("Authorization", "Bearer "+apiToken)
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return "I'm having trouble understanding right now. Try again later."
	// }
	// defer resp.Body.Close()

	// var response []HuggingFaceResponse
	// json.NewDecoder(resp.Body).Decode(&response)
	// if len(response) > 0 {
	// 	return response[0].GeneratedText
	// }
	// return "I'm not sure how to respond to that."
	return ""
}
