package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "¡Hola desde mi API desplegada con CI/CD y Floci!")
    })

    fmt.Println("Servidor corriendo en el puerto 8080...")
    http.ListenAndServe(":8080", nil)
}