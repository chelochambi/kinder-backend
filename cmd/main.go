package main

import (
	"kinder-backend/config"
	"kinder-backend/internal/handlers"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDatabase()

	// Configurar el router
	r := mux.NewRouter()

	// Rutas de los manejadores
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	// Iniciar el servidor
	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}
