# Etapa de compilación
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiamos AMBOS archivos de configuración de dependencias
COPY go.mod go.sum ./

# Descargamos dependencias validando las sumas
RUN go mod download

# Copiamos el código fuente de la aplicación
COPY main.go ./

# Compilamos el binario de Go de manera estática
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Etapa final de ejecución con una imagen ultraligera
FROM alpine:latest
WORKDIR /root/

# Copiamos el binario compilado desde la etapa anterior
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]