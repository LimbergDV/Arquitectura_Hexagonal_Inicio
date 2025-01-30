package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Conn_MySQL struct {
    DB *sql.DB
    Err string
}

func GetDBPool() *Conn_MySQL {
    error := ""
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    dbHost:= os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USERNAME")
    dbPass := os.Getenv("DB_PASSWORD")
    dbSchema := os.Getenv("DB_DATABASE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)

    db, err := sql.Open("mysql", dsn)

    if err != nil {
        error = fmt.Sprintf("error al abrir la base de datos: %w", err)
    }

    // Configuramos el pool de conexiones
    db.SetMaxOpenConns(10)

    // Para probar la conexión
    if err := db.Ping(); err != nil {
        db.Close()
        error = fmt.Sprintf("error al verificar la conexión de la base de datos %w", err)
    }

    fmt.Println("Conexión a la base de datos exitosamente")

    return &Conn_MySQL{DB: db, Err: error}


}
