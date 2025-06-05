package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chelochambi/kinder-backend/internal/db"
	"github.com/chelochambi/kinder-backend/internal/router"
	"github.com/joho/godotenv"
	"github.com/rs/cors" // <-- Agregamos esta importación
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	conn := db.GetDB()

	// Tu router personalizado
	r := router.NewRouter(conn)

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // permite tu frontend local
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor escuchando en el puerto %s\n", port)
	err = http.ListenAndServe(":"+port, corsHandler) // <-- usamos corsHandler
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
