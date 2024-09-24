package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"jobnbackpack.com/answer_generator/chat"
	"jobnbackpack.com/answer_generator/view"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	response := chat.CreateClient("Welcher Jünger ging auf Wasser?")

	var data []view.Choice
	err = json.Unmarshal([]byte(response), &data)

	fmt.Printf("%+v", data)

	p := tea.NewProgram(view.CreateView("Welcher Jünger ging auf Wasser?", data))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
