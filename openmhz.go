package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

type Calls struct {
	Calls []Call `json:"calls"`
}

type Call struct {
	ID              string    `json:"_id"`
	TalkGroupNumber int       `json:"talkgroupNum"`
	URL             string    `json:"url"`
	Filename        string    `json:"filename"`
	Timestamp       time.Time `json:"time"` //"time":"2024-11-21T22:22:09.000Z",
	Star            int       `json:"star"`
	Frequency       int64     `json:"freq"`
	Length          int       `json:"len"`
}

type ByTS []Call

func (a ByTS) Len() int           { return len(a) }
func (a ByTS) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTS) Less(i, j int) bool { return a[i].Timestamp.Before(a[j].Timestamp) }

func poll(channel string, params string, after time.Time) ([]Call, error) {
	tr := &http2.Transport{
		//MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	var req *http.Request
	var err error
	req, err = http.NewRequest("GET", fmt.Sprintf("https://api.openmhz.com/%s/calls?%s", channel, params), nil)
	if err != nil {
		return []Call{}, err
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("origin", "https://openmhz.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://openmhz.com/")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36'")

	if req.Body != nil {
		defer req.Body.Close()
	}

	resp, err := client.Do(req)
	if err != nil {
		return []Call{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Call{}, err
	}
	//log.Printf("body = %s", body)

	var c Calls
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Printf("ERR: Unmarshal: %s", err.Error())
		return []Call{}, err
	}

	//log.Printf("calls: %#v", c)

	finalC := make([]Call, 0)
	for _, v := range c.Calls {
		if v.Timestamp.Local().After(after.Local()) {
			log.Printf("DEBUG: %s is after %s", v.Timestamp.Local().String(), after.Local().String())
			finalC = append(finalC, v)
		}
	}

	log.Printf("INFO: Found %d calls after %s", len(finalC), after.Local().String())
	return finalC, nil
}

func getTempFile(url string) (string, error) {
	out, err := os.CreateTemp("/tmp", "audio")
	if err != nil {
		return "", err
	}

	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}

func tsFromCalls(c []Call) time.Time {
	ts := time.Now().Local()
	for _, v := range c {
		if v.Timestamp.Local().After(ts) {
			ts = v.Timestamp.Local()
		}
	}
	return ts
}

type talkgroupsObj struct {
	Talkgroups map[int]talkgroupObj `json:"talkgroups"`
}

type talkgroupObj struct {
	Alpha       string `json:"alpha"`
	Description string `json:"description"`
	Num         int    `json:"num"`
	ID          string `json:"_id"`
}

// Retrieve list of openmhz talkgroups for the specified channel
func talkgroups(channel string, fake bool) (map[int]string, error) {
	m := map[int]string{}

	body := []byte{}

	if !fake {
		tr := &http2.Transport{
			//MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}
		var req *http.Request
		var err error
		req, err = http.NewRequest("GET", fmt.Sprintf("https://api.openmhz.com/%s/talkgroups", channel), nil)
		if err != nil {
			return m, err
		}

		req.Header.Add("accept", "*/*")
		req.Header.Add("accept-language", "en-US,en;q=0.9")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("origin", "https://openmhz.com")
		req.Header.Add("pragma", "no-cache")
		req.Header.Add("priority", "u=1, i")
		req.Header.Add("sec-fetch-dest", "empty")
		req.Header.Add("sec-fetch-mode", "cors")
		req.Header.Add("sec-fetch-site", "same-site")
		req.Header.Add("referer", "https://openmhz.com/")
		req.Header.Add("sec-ch-ua", `"Chromium";v="142", "Google Chrome";v="142", "Not_A Brand";v="99"`)
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("sec-ch-ua-platform", "Linux")
		req.Header.Add("user-agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36`)

		if req.Body != nil {
			defer req.Body.Close()
		}

		resp, err := client.Do(req)
		if err != nil {
			return m, err
		}
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return m, err
		}
		//log.Printf("body = %s", body)
	} else {
		var err error
		body, err = os.ReadFile("talkgroups.json")
		if err != nil {
			return m, err
		}
	}

	var c talkgroupsObj
	err := json.Unmarshal(body, &c)
	if err != nil {
		log.Printf("ERR: Unmarshal: %s", err.Error())
		log.Printf("body = %s", body)
		return m, err
	}

	for _, v := range c.Talkgroups {
		m[v.Num] = v.Description
	}

	//log.Printf("calls: %#v", c)

	return m, nil
}
