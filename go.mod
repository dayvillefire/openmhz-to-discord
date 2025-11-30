module github.com/dayvillefire/openmhz-to-discord

go 1.24.2

replace (
	github.com/dayvillefire/dgvoice => ../dgvoice
	github.com/dayvillefire/dgvoice/mpg123 => ../dgvoice/mpg123
)

require (
	github.com/alexsasharegan/dotenv v0.0.0-20171113213728-090a4d1b5d42
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dayvillefire/dgvoice v0.0.0-20251130190534-5aea918fd802
	golang.org/x/net v0.47.0
)

require (
	github.com/Carmen-Shannon/gopus v1.0.0 // indirect
	github.com/dayvillefire/dgvoice/mpg123 v0.0.0-20251105145907-bfc8f72e3c4d // indirect
	github.com/ebitengine/oto/v3 v3.4.0 // indirect
	github.com/ebitengine/purego v0.9.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hajimehoshi/ebiten/v2 v2.9.4 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
)
