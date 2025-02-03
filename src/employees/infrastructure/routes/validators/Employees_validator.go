package validators

import (
	"errors"
	"introduccion_go/src/employees/domain"
)


func CheckEmployee (employee domain.Employee) error {

	if employee.Id < 0{
		return errors.New("El id tiene que ser mayor a 0")
	}

	if employee.First_name == ""{
		return errors.New("No puedes poner un nombre vacio")
	}

	if employee.Last_name == ""{
		return errors.New("No puedes poner un apellido vacio")
	}

	if employee.Age < 0 {
		return errors.New("No puedes poner una edad menor a 0")
	}

	if employee.Age < 18 {
		return errors.New("No puedes poner un menor de edad a trabajar")
	}

	if employee.Curp == "" {
		return errors.New("La curp no puede estar vacia")
	}

	if employee.Phone_number == "" {
		return errors.New("No puedes dejar el campo de número vacío")
	}

	if employee.Salary < 0 {
		return errors.New("No puedes poner un salario menor a 0")
	}

	return nil
}