package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"jobnbackpack.com/answer_generator/models"
)

const (
	api_endpoint = "https://api.openai.com/v1/chat/completions"
)

func AskGPT(question string) []models.Choice {
	response := createClient(question)

	var data []models.Choice
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Printf("%v", err)
	}

	return data
}

func createClient(question string) string {
	apiKey := os.Getenv("GPT_API_TOKEN")
	client := resty.New()

	prompt := "I want to create a quiz with biblical questions. I want you to give me four alternative answers to following question: " + question + " The format for the answers should be like this {\"value\": string; \"correct\": boolean}[]"

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("OpenAI-Organization", os.Getenv("GPT_ORG_ID")).
		SetHeader("OpenAI-Project", os.Getenv("GPT_PROJECT_ID")).
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "user", "content": prompt}},
			"max_tokens": 100,
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
		return ""
	}

	// Extract the content from the JSON response
	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	fmt.Println(content)
	return content
}
