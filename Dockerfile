# Fliboobstier bot container
FROM golang:alpine
WORKDIR /go/src/fliboobstier
COPY fliboobstier.go .
RUN apk add --no-cache git \
    && go get github.com/go-telegram-bot-api/telegram-bot-api \
    && go get -d -v ./... \
    && go install -v ./...
ENTRYPOINT ["/go/bin/fliboobstier"]