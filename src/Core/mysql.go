package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


type Conn_MySQL struct {
   DB*sql.DB
   Err string
}


func GetPool() *Conn_MySQL{
    error := ""
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env %v", err)
    }

    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USERNAME")
    dbPass := os.Getenv("DB_PASSWORD")
    dbDataBase := os.Getenv("DB_DATABASE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbHost, dbUser, dbPass, dbDataBase)

    db, err:= sql.Open("mysql", dsn)

    if err != nil {
        error = fmt.Sprintf("Error al abrir la base de datos %w", err)
    }
    
    db.SetMaxOpenConns(10)

    if err := db.Ping(); err != nil {
        db.Close()
        error = fmt.Sprintf("error al verificar la conexión a la base de datos: %w", err)
    }
    return &Conn_MySQL{DB: db, Err: error}
}

func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) (*sql.Rows) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf ("error al ejecutar la consulta SELECT: %w", err)
	}

	return rows
}
