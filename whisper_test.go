package main

import (
	"log"
	"os"
	"testing"

	"github.com/alexsasharegan/dotenv"
)

func Test_WhisperPost(t *testing.T) {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	whisperServer := os.Getenv("WHISPER_SERVER_URL")

	// Example usage of DiscordVoice
	s, err := postWhisper(whisperServer, "pfield-1401-1750268957.mp3")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
