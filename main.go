package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/alexsasharegan/dotenv"
)

var (
	token           = flag.String("token", "", "Discord API token")
	guild           = flag.String("guild", "", "Guild ID")
	channel         = flag.String("channel", "", "Channel ID")
	openmhzChannel  = flag.String("omhz", "", "OpenMHZ Channel ID")
	pollingInterval = flag.Int("poll-interval", 0, "Polling interval")
)

func main() {
	flag.Parse()
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if *token == "" {
		*token = os.Getenv("DISCORD_TOKEN")
	}
	if *guild == "" {
		*guild = os.Getenv("DISCORD_GUILD")
	}
	if *channel == "" {
		*channel = os.Getenv("DISCORD_CHANNEL")
	}
	if *openmhzChannel == "" {
		*openmhzChannel = os.Getenv("OPENMHZ_CHANNEL")
	}
	if *pollingInterval == 0 {
		*pollingInterval, _ = strconv.Atoi(os.Getenv("OPENMHZ_POLLING_INTERVAL"))
	}

	ds := DiscordVoice{}
	log.Printf("INFO: Init with token")
	err = ds.Init(*token, *guild, *channel)
	if err != nil {
		panic(err)
	}

	// Initial poll
	initialTime := time.Now().Local()
	log.Printf("INFO: Starting polling after %s", initialTime)
	calls, err := poll(*openmhzChannel, initialTime)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return
	}

	// Starting timestamp
	ts := tsFromCalls(calls)
	done := false
	log.Printf("INFO: Starting on calls after %s", ts.Local().String())

	go func() {
		log.Printf("INFO: Entering playback loop")
		for {
			consumeQueue(func(c Call) {
				fn, err := getTempFile(c.URL)
				if err != nil {
					log.Printf("ERR: getTempFile: %s", err.Error())
					return
				}
				log.Printf("INFO: Play %s", fn)
				ds.Play(fn)
				log.Printf("INFO: Sleeping for duration of file %d seconds", c.Length)
				time.Sleep(time.Duration(c.Length) * time.Second)
				os.Remove(fn)
			})

			if done {
				log.Printf("INFO: Exiting play loop")
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	log.Printf("INFO: Beginning poll loop")
	for !done {
		time.Sleep(time.Second * time.Duration(*pollingInterval))
		// Poll for calls
		calls, err = poll(*openmhzChannel, ts.Local())
		if len(calls) == 0 {
			log.Printf("INFO: No calls, going back to wait loop")
			continue
		}
		// Update timestamp
		ts = tsFromCalls(calls)
		// Sort calls
		sort.Sort(ByTS(calls))
		// Play calls in channel
		for _, v := range calls {
			enqueueItem(v)
		}
	}

	// Close connections
	ds.Close()
}
