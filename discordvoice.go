package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dayvillefire/dgvoice"
)

type DiscordVoice struct {
	discordSession *discordgo.Session
	discordInit    bool
	dgv            *discordgo.VoiceConnection
	internal       bool
}

func (d *DiscordVoice) Init(token, guild, channel string, internal bool) error {
	var err error
	if d.discordInit {
		return fmt.Errorf("ERR: already intiialized: %w", err)
	}

	d.internal = internal

	d.discordSession, err = discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("ERR: New(): %w", err)
	}

	d.discordSession.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)

	err = d.discordSession.Open()
	if err != nil {
		return fmt.Errorf("ERR: Open(): %w", err)
	}

	log.Printf("INFO: ChannelVoiceJoin")
	d.dgv, err = d.discordSession.ChannelVoiceJoin(guild, channel, false, true)
	if err != nil {
		log.Printf("ERR: Connecting to voice channel")
		return err
	}

	d.discordInit = true
	return nil
}

func (d *DiscordVoice) Play(filepath string) {
	ch := make(chan bool)
	if d.internal {
		log.Printf("INFO: Playing audio file internally: %s", filepath)
		go dgvoice.PlayAudioFileInternal(d.dgv, filepath, ch)
	} else {
		log.Printf("INFO: Playing audio file externally: %s", filepath)
		go dgvoice.PlayAudioFile(d.dgv, filepath, ch)
	}
	log.Printf("INFO: Waiting for stop signal")
	<-ch
}

func (d *DiscordVoice) Close() {
	d.dgv.Close()
	d.discordSession.Close()
}
