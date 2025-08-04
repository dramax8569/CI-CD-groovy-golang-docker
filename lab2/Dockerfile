# Etapa 1: build estático para Linux
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

# IMPORTANTE: build completamente estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Etapa 2: imagen mínima compatible
FROM scratch

WORKDIR /app
COPY --from=builder /app/app .

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["./app"]
