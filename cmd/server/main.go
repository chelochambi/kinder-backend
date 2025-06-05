package main

import (
	"log"
	"net/http"
	"os"

	/* "fmt" */
	"github.com/chelochambi/kinder-backend/internal/db"
	"github.com/chelochambi/kinder-backend/internal/router"
	"github.com/joho/godotenv"
	/* "golang.org/x/crypto/bcrypt" */)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	/* hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash)) */

	conn := db.GetDB()

	r := router.NewRouter(conn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor escuchando en el puerto %s\n", port)
	http.ListenAndServe(":"+port, r)

}
