package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	// Acceder a las variables de entorno
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Cadena de conexión a PostgreSQL
	// Para entorno de desarrollo
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Conectar a la base de datos
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	} else {
		log.Println("Conectado a la base de datos")
	}
}
