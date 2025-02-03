package application

import "introduccion_go/src/employees/domain"

type GetAllEmployees struct {
	db domain.IEmployee
}

func NewGetAllEmployees(db domain.IEmployee) *GetAllEmployees {
	return &GetAllEmployees{db: db}
}

func (le *GetAllEmployees) Run () []domain.Employee {
	return le.db.GetAll()
}