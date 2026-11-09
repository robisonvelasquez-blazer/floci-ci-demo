# Usamos la imagen oficial de Go para compilar
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiamos primero los archivos de dependencias y código fuente
COPY go.mod ./
COPY main.go ./

# Compilamos la aplicación
RUN go build -o main main.go

# Imagen final ligera para ejecución
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]