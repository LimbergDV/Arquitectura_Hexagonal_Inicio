package infraestructure

import (
	"database/sql"
	"fmt"
	"introduccion_go/src/products/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db:db}
}

func (mysql *MySQL) Save(product domain.Product) error {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"
	_, err := mysql.db.Exec(query, product.GetName(), product.GetPrice())
	
	if err != nil {
		return err
	}

	fmt.Println("Producto creado")
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Product, error){
	query := "SELECT * FROM product"
	rows, err := mysql.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []domain.Product

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

	fmt.Println("Lista de productos")

	return products, nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ? "
	_, err := mysql.db.Exec(query, id)
	fmt.Println("PRODUCTO ELIMINADO")
	return err
}

func (mysql *MySQL) Update(id int, product domain.Product) error {
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	_, err := mysql.db.Exec(query, product.GetName(), product.GetPrice(), id)
	fmt.Println("PRODUCTO ACTUALIZADO")
	return err
}