package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"jobnbackpack.com/answer_generator/logger"
	"jobnbackpack.com/answer_generator/view"
)

func main() {
	initEnv()
	f := initLogger()
	defer f.Close()

	p := tea.NewProgram(view.CreateView())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func initLogger() *os.File {
	f, err := logger.LogToFile("debug.log", "DEBUG")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	log.Print("I started..")

	return f
}
