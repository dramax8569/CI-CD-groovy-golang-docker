FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Binario estático compatible con Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 9090

CMD ["./app"]
