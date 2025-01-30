package infraestructure

import (
	"database/sql"
	"fmt"
	"introduccion_go/src/Core"
	"introduccion_go/src/products/domain"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(product domain.Product) error {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"
	_, err := mysql.
	
	if err != nil {
		return err
	}

	fmt.Println("Producto creado")
	return nil
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