package main

import (
	"log"
	"os"
	"testing"

	"github.com/alexsasharegan/dotenv"
)

func Test_DiscordVoice(t *testing.T) {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	guild := os.Getenv("DISCORD_GUILD")
	channel := os.Getenv("DISCORD_CHANNEL")

	// Example usage of DiscordVoice
	d := &DiscordVoice{}
	err = d.Init(token, guild, channel, false)
	if err != nil {
		t.Fatalf("Failed to initialize Discord voice: %v", err)
	}
	d.Play("pfield-1401-1750268957.mp3")
	d.Close()
}

func Test_DiscordVoiceInternal(t *testing.T) {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	guild := os.Getenv("DISCORD_GUILD")
	channel := os.Getenv("DISCORD_CHANNEL")

	// Example usage of DiscordVoice
	d := &DiscordVoice{}
	err = d.Init(token, guild, channel, true)
	if err != nil {
		t.Fatalf("Failed to initialize Discord voice: %v", err)
	}
	d.Play("pfield-1401-1750268957.mp3")
	d.Close()
}
