FROM golang:1.19.2-alpine3.16 AS builder 

COPY . /github.com/dazai404/pocketerist-bot
WORKDIR /github.com/dazai404/pocketerist-bot

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest 

WORKDIR /root/

COPY --from=0 /github.com/dazai404/pocketerist-bot/bin/bot .
COPY --from=0 /github.com/dazai404/pocketerist-bot/configs configs/

EXPOSE 80

CMD ["./bot"]
