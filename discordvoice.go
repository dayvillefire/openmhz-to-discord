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
}

func (d *DiscordVoice) Init(token, guild, channel string) error {
	var err error
	if d.discordInit {
		return fmt.Errorf("ERR: already intiialized: %w", err)
	}

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
	dgvoice.PlayAudioFile(d.dgv, filepath, ch)
	<-ch
}

func (d *DiscordVoice) Close() {
	d.dgv.Close()
	d.discordSession.Close()
}
