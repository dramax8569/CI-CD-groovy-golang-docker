# Microservicio en Go + CI/CD

Este es un microservicio HTTP básico hecho en Go que expone el endpoint `/ping`.

## Estructura

- `main.go`: Entry point
- `handler/`: Contiene la lógica de negocio y pruebas
- `Dockerfile`: Imagen multistage optimizada
- `Jenkinsfile`: Pipeline CI/CD con etapas build → test → docker build → push

## Ejecutar localmente

```bash
go build -o app
./app
curl http://localhost:8080/ping
