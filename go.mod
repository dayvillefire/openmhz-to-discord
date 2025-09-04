module github.com/dayvillefire/openmhz-to-discord

go 1.24.2

replace (
	github.com/dayvillefire/dgvoice => ../dgvoice
	github.com/dayvillefire/dgvoice/mpg123 => ../dgvoice/mpg123
)

require (
	github.com/alexsasharegan/dotenv v0.0.0-20171113213728-090a4d1b5d42
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dayvillefire/dgvoice v0.0.0-20250904121917-21df3be04847
	golang.org/x/net v0.43.0
)

require (
	github.com/Carmen-Shannon/gopus v1.0.0 // indirect
	github.com/dayvillefire/dgvoice/mpg123 v0.0.0-20250827184712-8ccad88ca83b // indirect
	github.com/ebitengine/oto/v3 v3.3.3 // indirect
	github.com/ebitengine/purego v0.8.4 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hajimehoshi/ebiten/v2 v2.8.8 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	layeh.com/gopus v0.0.0-20210501142526-1ee02d434e32 // indirect
)
