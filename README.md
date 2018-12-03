# Fliboobstier

### Envs

**FLIBOOBSTIER_TG_TOKEN** - set Telegram bot token   
`export FLIBOOBSTIER_TG_TOKEN="some_very_secret_token"`

**HTTP_PROXY** - use proxy (optional)   
`export HTTP_PROXY="socks5://login:passwd@example.com:1080"`

#### Manual build and run

```
make deps && make && make install
$GOPATH/bin/fliboobstier
```

### Docker build and run
```
docker build -t fliboobstier .

docker run -d \
-e HTTP_PROXY=$HTTP_PROXY \
-e FLIBOOBSTIER_TG_TOKEN=$FLIBOOBSTIER_TG_TOKEN \
fliboobstier
```

License
----
Beer and pizza license
