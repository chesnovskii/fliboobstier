# Fliboobstier

### Environment

**`TG_TOKEN`** - Telegram bot token

For example:

    $ export TG_TOKEN="some_very_secret_token"
    $ echo $TG_TOKEN
    some_very_secret_token

**`HTTP_PROXY`** - Use proxy (if needed)

For example:

    $ export HTTP_PROXY="socks5://login:passwd@example.com:1080"
    $ echo $HTTP_PROXY
    socks5://login:passwd@example.com:1080

#### Building for source

```
$ go get github.com/go-telegram-bot-api/telegram-bot-api
$ go build -o fliboobstier .
$ go run fliboobstier
```

### Docker
```
$ docker build -t fliboobstier .
```

```
docker run -d \
-e HTTP_PROXY=$HTTP_PROXY \
-e TG_TOKEN=$SEAL_TG_TOKEN \
fliboobstier
```

License
----
Beer and pizza license