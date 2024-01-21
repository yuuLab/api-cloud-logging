FROM golang:1.21.6-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go


# Multi-Stage build 
FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]