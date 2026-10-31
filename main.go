package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	// Configure bot settings using an environment variable for your Telegram Bot Token
	pref := tele.Settings{
		Token:  os.Getenv("TELE_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Please check your TELE_TOKEN environment variable: %v", err)
		return
	}

	// Define a command handler (e.g., when a user sends /hello)
	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	// Start the bot
	b.Start()
}