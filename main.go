package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	api_endpoint = "https://api.openai.com/v1/chat/completions"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("GPT_API_TOKEN")
	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("OpenAI-Organization", os.Getenv("GPT_ORG_ID")).
		SetHeader("OpenAI-Project", os.Getenv("GPT_PROJECT_ID")).
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "user", "content": "I want to create a quiz with biblical questions. I want you to give me three alternative answers to following question: Which deciple walked on water?"}},
			"max_tokens": 50,
		}).
		Post(api_endpoint)
	if err != nil {
		log.Fatalf("Error while sending send the request: %v", err)
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error while decoding JSON response:", err)
		return
	}

	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(content)
}
