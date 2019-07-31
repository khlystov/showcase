FROM golang:1.12.7-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest


WORKDIR /app
COPY --from=builder /app .

EXPOSE 3200

ENTRYPOINT ["./main"]