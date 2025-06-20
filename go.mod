module github.com/dayvillefire/openmhz-to-discord

go 1.24.2

replace (
	github.com/dayvillefire/dgvoice => ../dgvoice
	github.com/dayvillefire/dgvoice/mpg123 => ../dgvoice/mpg123
)

require (
	github.com/alexsasharegan/dotenv v0.0.0-20171113213728-090a4d1b5d42
	github.com/bwmarrin/discordgo v0.29.0
	github.com/dayvillefire/dgvoice v0.0.0-20250320123128-426d91671cf2
	golang.org/x/net v0.41.0
)

require (
	github.com/bobertlo/go-mpg123 v0.0.0-20211210004329-c83f21a0fd39 // indirect
	github.com/dayvillefire/dgvoice/mpg123 v0.0.0-20250619171946-f4f198d25038 // indirect
	github.com/ebitengine/oto/v3 v3.3.3 // indirect
	github.com/ebitengine/purego v0.8.4 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hajimehoshi/ebiten/v2 v2.8.8 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	layeh.com/gopus v0.0.0-20210501142526-1ee02d434e32 // indirect
)
