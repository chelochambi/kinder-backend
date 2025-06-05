package main

import (
    "log"
    "net/http"
    "os"

    "github.com/chelochambi/kinder-backend/internal/db"
    "github.com/chelochambi/kinder-backend/internal/router"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No se pudo cargar el archivo .env")
    }

    conn := db.GetDB()

    r := router.NewRouter(conn)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Servidor escuchando en el puerto %s\n", port)
    http.ListenAndServe(":"+port, r)
}
