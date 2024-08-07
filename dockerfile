FROM golang:1.18 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .

FROM alpine:latest

RUN apk --no-cache add tzdata

COPY --from=builder /app/main.app /usr/local/bin

EXPOSE 8000

CMD ["main.app"]