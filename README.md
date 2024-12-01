# Telegram songs dl
I'm starting to mix music and I would like to "automize" songs downloading in a simple way.
This is a simple bot to share youtube urls and save my favourite songs.

At this moment it just saves the urls in a file and then with a script I download the songs
that are listed in that file. In a future I would like to automatically download the songs.

## Use the bot
**Using docker**
This image is compatible with `amd64` and `arm64` architectures.
```yaml
services:
  bot:
    image: 4strodev/telegram-songs-dl:latest
    restart: always
    env_file: .env
  ## If you prefer you can set the environment variables directly
  ## on the docker compose environment entry
    #environment:
      #TELEGRAM_TOKEN: "<your token>"
      #SONGS_FILE: "<songs file path>"
```

# Contribute
If you want to contribute to this project or fork it to create your own version you are free to do that.

## Prerequisites
To work confortably with this project I recommend to you to use this
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
