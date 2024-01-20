FROM golang:1.21.6-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api ./cmd/api/main.go


# Multi-Stage build 
FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/api .

EXPOSE $POST

CMD ["./api"]