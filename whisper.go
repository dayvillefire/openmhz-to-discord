package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

// Provides whisper-server integration
func whisper(c Call, filepath string) {
	log.Printf("INFO: Starting transcribe to %s", *whisperServerUrl)
	defer os.Remove(filepath)

	// Transcribe
	txt, err := postWhisper(*whisperServerUrl, filepath)
	if err != nil {
		log.Printf("ERR: whisper: %s", err.Error())
		return
	}

	// Resolve talkgroup
	t, found := tg[c.TalkGroupNumber]
	if !found {
		t = fmt.Sprintf("Talkgroup #%d", c.TalkGroupNumber)
	}

	// Post text to discord

	fp, err := os.Open(filepath)
	if err != nil {
		log.Printf("ERR: whisper: error opening file: %s", err.Error())
		_, err = ds.discordSession.ChannelMessageSend(
			*channelTranscribe,
			fmt.Sprintf("%s: %s: %s [audio](%s)",
				c.Timestamp.Local().Format("2006-01-02 15:04:05"),
				t,
				txt,
				c.URL,
			))
		if err != nil {
			log.Printf("ERR: whisper: SendMessage(): %s", err.Error())
			return
		}
		return
	}
	defer fp.Close()
	msg := &discordgo.MessageSend{
		Content: fmt.Sprintf("%s: %s: %s [audio](%s)",
			c.Timestamp.Local().Format("2006-01-02 15:04:05"),
			t,
			txt,
			c.URL,
		),
		Files: []*discordgo.File{
			{
				Name:        c.Filename,
				ContentType: "audio/mp3",
				Reader:      fp,
			},
		},
	}

	_, err = ds.discordSession.ChannelMessageSendComplex(*channelTranscribe, msg)
	if err != nil {
		log.Printf("ERR: whisper: SendMessage(): %s", err.Error())
		return
	}
}

func postWhisper(whisperServerUrl string, filepath string) (string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("ERR: postWhisper: os.Open: %s", err)
		return "", err
	}

	part, err := w.CreateFormFile("audio_file", file.Name())
	if err != nil {
		log.Printf("ERR: postWhisper: multipart.CreateFormFile: %s", err)
		return "", err
	}
	b, err := io.ReadAll(file)
	if err != nil {
		log.Printf("ERR: postWhisper: io.ReadAll: %s", err)
		return "", err
	}
	part.Write(b)
	w.Close()

	client := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest("POST",
		fmt.Sprintf("%s/asr?encode=true&task=transcribe&language=en&vad_filter=true&word_timestamps=false&output=txt",
			whisperServerUrl),
		&buf)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", w.FormDataContentType())

	if req.Body != nil {
		defer req.Body.Close()
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	return nil
}
