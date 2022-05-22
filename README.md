# Telegram echo bot

This bot can send you notification from any computer with curl installed.

## Usage

````
# build executable
go build .

# run server
./techo-go -port <port to serve> -owner_id <your telegram id> -bot_token <token of your bot>
````

To send yourself a notification, just run
````
curl http://<server_host>:<server_port>/echo/message
````
