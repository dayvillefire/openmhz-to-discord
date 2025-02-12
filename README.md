# OPENMHZ-TO-DISCORD

Daemon that streams OpenMHZ radio streams to a Discord voice channel.

## Install and run via docker

### `docker-compose.yml` ###

```
version: '3.5'

services:
  openmhz-to-discord:
    container_name: openmhz-to-discord
    image: jbuchbinder/openmhz-to-discord
    restart: always
    volumes:
      - .env:/.env:ro
    env_file:
      - .env
```

### `.env` ###

```
DISCORD_TOKEN="MTAxxxxxxxxxxxxxxxxxxxxxxx.xxxxxx.xxxxxxxxxxxxxxxxxxxxxx_xxxxxxxxxx_xxxx" # Token for discord app
DISCORD_GUILD="1000000000000000000"  # Discord server ID
DISCORD_CHANNEL="1000000000000000000" # Discord channel ID
OPENMHZ_CHANNEL="ctfire" # name of the openmhz channel
OPENMHZ_POLLING_INTERVAL=3
```

### Starting ###

```
docker-compose up -d
```

## Caveats ##

* Due to the way discord functions, a single app ID (application) can be connected to a single channel. You cannot run multiple instances of this docker image with the same `DISCORD_TOKEN`, otherwise earlier channels will be disconnected in favor of later channels.

