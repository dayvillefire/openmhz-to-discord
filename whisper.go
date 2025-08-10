package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// Provides whisper-server integration
func whisper(c Call, filepath string) {
	log.Printf("INFO: Starting transcribe to %s", *whisperServerUrl)

	// Transcribe
	txt, err := postWhisper(*whisperServerUrl, filepath)
	if err != nil {
		log.Printf("ERR: whisper: %s", err.Error())
		return
	}

	// Post text to discord
	msg := fmt.Sprintf("%.3f: %s", float64(c.Frequency)/1000000, txt)
	_, err = ds.discordSession.ChannelMessageSend(*channelTranscribe, msg)
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
