package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"jobnbackpack.com/answer_generator/chat"
	"jobnbackpack.com/answer_generator/logger"
	"jobnbackpack.com/answer_generator/view"
)

func main() {
	initEnv()
	initLogger()

	p := tea.NewProgram(view.CreateView("Welcher Jünger ging auf Wasser?", askGPT()))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func askGPT() []view.Choice {
	response := chat.CreateClient("Welcher Jünger ging auf Wasser?")

	var data []view.Choice
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Printf("%v", err)
	}

	return data
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func initLogger() {
	f, err := logger.LogToFile("debug.log", "DEBUG")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
}
