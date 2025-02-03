package application

import "introduccion_go/src/employees/domain"





type CreateEmployee struct{
	db domain.IEmployee
}

func NewCreateEmployee (db domain.IEmployee) *CreateEmployee {
	return &CreateEmployee{db: db}
}

func (ce *CreateEmployee) Run (employee domain.Employee) (uint, error) {
	return ce.db.Save(employee)
}