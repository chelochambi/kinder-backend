package db

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/jackc/pgx/v5/stdlib" // driver pgx para database/sql
)

var DB *sql.DB

func InitDB() *sql.DB {
    if DB != nil {
        return DB
    }

    dsn := os.Getenv("DATABASE_URL")
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("Error al abrir base de datos: %v", err)
    }

    // Opcional: verificar conexión
    if err := db.Ping(); err != nil {
        log.Fatalf("Error al conectar a base de datos: %v", err)
    }

    DB = db
    return DB
}

func GetDB() *sql.DB {
    if DB == nil {
        return InitDB()
    }
    return DB
}
