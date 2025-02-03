package infrastructure

import (
	"fmt"
	core "introduccion_go/src/Core"
	"introduccion_go/src/employees/domain"
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

func (mysql *MySQL) Save(employee domain.Employee) (uint, error) {
	query := "INSERT INTO employees (first_name, last_name, age, curp, phone_number, salary) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := mysql.conn.ExecutePreparedQuery(query, employee.First_name, employee.Last_name, employee.Age, employee.Curp, employee.Phone_number, employee.Salary)
	if err != nil {
		fmt.Println("Error al preparar la consulta: %v", err)
		return 0, err
	}
	id, _ := res.LastInsertId()
	fmt.Println("Empleado creado")
	return uint(id), nil
}

func (mysql *MySQL) GetAll() []domain.Employee {
	query := "SELECT * FROM employees"
	var employees []domain.Employee
	rows := mysql.conn.FetchRows(query)
	
	if rows == nil {
		fmt.Println("No se obtuvieron los datos")
		return employees
	}
	defer rows.Close()

	for rows.Next() {
		var e domain.Employee
		if err := rows.Scan(&e.Id, &e.First_name, &e.Last_name, &e.Age, &e.Curp, &e.Phone_number, &e.Salary); err != nil {
			fmt.Println("Error al escanear la fila: %w", err)
		} else {
			employees = append(employees, e)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterando sobre las filas: %w", err)
	}

	fmt.Println("Lista de empleados")
	return employees
}

func (mysql *MySQL) Delete(id int) (uint, error) {
	query := "DELETE FROM employees WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	return uint(rows), nil
}

func (mysql *MySQL) Update(id int, employee domain.Employee) (uint, error) {
	query := "UPDATE employees SET first_name = ?, last_name = ?, age = ?, curp = ?, phone_number = ?, salary = ? WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, employee.First_name, employee.Last_name, employee.Age, employee.Curp, employee.Phone_number, employee.Salary, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	return uint(rows), nil
}
