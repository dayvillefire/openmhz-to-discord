package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

// Provides whisper-server integration
func whisper(c Call, filepath string) {
	log.Printf("INFO: Starting transcribe to %s", *whisperServerUrl)

	// Transcribe
	txt, err := postWhisper(filepath)
	if err != nil {
		log.Printf("ERR: whisper: %s", err.Error())
		return
	}

	// Post text to discord
	msg := fmt.Sprintf("%s", txt)
	_, err = ds.discordSession.ChannelMessageSend(*channelTranscribe, msg)
	if err != nil {
		log.Printf("ERR: whisper: SendMessage(): %s", err.Error())
		return
	}

}

func postWhisper(filepath string) (string, error) {
	tr := &http2.Transport{
		//MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	b, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	client := &http.Client{Transport: tr}
	var req *http.Request
	req, err = http.NewRequest("POST",
		fmt.Sprintf("%s/asr?encode=true&task=transcribe&language=en&vad_filter=true&word_timestamps=false&output=txt",
			*whisperServerUrl),
		bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "multipart/form-data")

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
