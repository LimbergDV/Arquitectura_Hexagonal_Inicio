package infraestructure

import (
	"fmt"
	"introduccion_go/src/Core"
	"log"
)

type MySQL struct {
	conn *core.ConnectToDB
}

func NewMySQL() *MySQL {
	conn := core.ConnectToDB()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save() {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"
	name := "Big-cola"
	price := 12.50

	result, err := mysql.conn.ExecutePreparedQuery(query, name, price)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
}

func (mysql *MySQL) GetAll() {
	query := "SELECT * FROM product"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()
	for rows.Next() {
		var idproduct int
		var name string
		var price float32
		if err := rows.Scan(&idproduct, &name, &price); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		fmt.Printf("ID: %d, Name: %s, Price: %f\n", idproduct, name, price)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("error iterando sobre las filas: %w", err)
	}
}