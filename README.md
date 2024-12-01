# Telegram songs dl
I'm starting to mix music and I would like to "automize" songs downloading in a simple way.
This is a simple bot to share youtube urls and save my favourite songs.

At this moment it just saves the urls in a file and then with a script I download the songs
that are listed in that file.

## Prerequisites

**Required**
- Docker and docker compose

**Optional**
- [Task](https://taskfile.dev)

## Start bot
Setup environment variables

### Using a .env file
Put a `.env` file in the bot folder with the following content
```
TELEGRAM_TOKEN="<Your telegram bot token>"
SONGS_FILE="./songs.txt" # Change this if you want
```

### Setting env variables on docker compose file
See the docker compose file to setup env variables

### Start docker container
**Using task**

    task bot:start

**Using docker compose**

    docker compose -f bot/docker-compose.yml up
