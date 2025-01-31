package infraestructure

import (
	"fmt"
	core "introduccion_go/src/Core"
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

func (mysql *MySQL) Save(product domain.Product) (uint, error) {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"
	res, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price)

	if err != nil {
		fmt.Println("Error al preparar la consulta: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId()

	fmt.Println("Producto creado")
	return uint(id), nil
}

func (mysql *MySQL) GetAll() []domain.Product {
	query := "SELECT * FROM product"

	var products [] domain.Product
	rows := mysql.conn.FetchRows(query)

	if rows == nil {
		fmt.Println("No se obtuvieron los datos")
		return products
	}

	defer rows.Close()

	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.Id, p.Name, p.Price); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}

	}

	if err := rows.Err(); err != nil {
		fmt.Println("error iterando sobre las filas: %w", err)
	}

	fmt.Println("Lista de productos")

	return products
}

func (mysql *MySQL) Delete(id int) (uint, error) {
	query := "DELETE FROM products WHERE id = ? "

	res, err := mysql.conn.ExecutePreparedQuery(query, id)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta %v", err)
		return 0, err
	}
	
	rows, _ := res.RowsAffected()

	return uint(rows), nil
}

func (mysql *MySQL) Update(id int, product domain.Product) (uint, error) {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price)
	
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
	}

	rows, _ := res.RowsAffected()

	return uint(rows), nil
}
