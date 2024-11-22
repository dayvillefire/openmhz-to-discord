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
	calls, err := poll(*openmhzChannel, time.Now().Add(time.Hour))
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return
	}

	// Starting timestamp
	ts := tsFromCalls(calls)
	done := false

	for !done {
		time.Sleep(time.Second * time.Duration(*pollingInterval))
		// Poll for calls
		calls, err = poll(*openmhzChannel, ts)
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
			fn, err := getTempFile(v.URL)
			if err != nil {
				log.Printf("ERR: getTempFile: %s", err.Error())
				continue
			}

			// TODO: FIXME: IMPLEMENT: XXX: Move this to another thread and queue items so as not to block this thread

			log.Printf("INFO: Play %s", fn)
			ds.Play(fn)
			log.Printf("INFO: Sleeping for duration of file %d seconds", v.Length)
			time.Sleep(time.Duration(v.Length) * time.Second)
			os.Remove(fn)
		}
		//log.Printf("%#v", calls)
	}

	// Close connections
	ds.Close()
}
