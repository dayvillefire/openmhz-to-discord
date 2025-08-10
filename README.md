# OPENMHZ-TO-DISCORD

Daemon that streams [OpenMHZ](https://openmhz.com/) radio streams to a Discord voice channel. Also optionally allows transcription via [whisper-server](https://github.com/nalbion/whisper-server), using OpenAI's [whisper](https://github.com/openai/whisper) as a backend, to a separate Discord text channel.

## Install and run via docker

### `docker-compose.yml` ###

Make sure to install `libsound` and `libmpg123-0` if you're using a debian variant.

```
version: '3.5'

services:
  openmhz-to-discord:
    container_name: openmhz-to-discord
    image: jbuchbinder/openmhz-to-discord
    restart: unless-stopped
    volumes:
      - .env:/.env:ro
      - /usr/lib/x86_64-linux-gnu/libasound.so.2:/usr/lib/libasound.so.2:ro
      - /usr/lib/x86_64-linux-gnu/libmpg123.so.0:/usr/lib/libmpg123.so.0:ro
    env_file:
      - .env
    depends_on:
      - whisper-server

  whisper-server:
    container_name: whisper-server
    image: onerahmet/openai-whisper-asr-webservice:latest
    restart: unless-stopped
    environment:
      - ASR_MODEL=small
      - ASR_ENGINE=faster_whisper
      - ASR_DEVICE=cpu
    volumes:
      - ./cache:/root/.cache
    ports:
      - 10300:9000
```

### `.env` ###

```
DISCORD_TOKEN="MTAxxxxxxxxxxxxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxx_xxxx" # Token for discord app
DISCORD_GUILD="1000000000000000000"  # Discord server ID
DISCORD_CHANNEL="1000000000000000000" # Discord channel ID
OPENMHZ_CHANNEL="ctfire" # name of the openmhz channel
OPENMHZ_POLLING_INTERVAL=7
WHISPER_SERVER_URL="http://whisper-server:10300"
WHISPER_DISCORD_CHANNEL="100000000000000000" # scanner text output channel
```

### Starting ###

```
docker-compose up -d
```

## Caveats ##

* Due to the way discord functions, a single app ID (application) can be connected to a single channel. You cannot run multiple instances of this docker image with the same `DISCORD_TOKEN`, otherwise earlier channels will be disconnected in favor of later channels.

