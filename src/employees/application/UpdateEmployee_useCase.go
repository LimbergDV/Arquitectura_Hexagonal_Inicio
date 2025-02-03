package application

import "introduccion_go/src/employees/domain"

type UpdateEmployee struct {
	db domain.IEmployee
}

func NewUpdateEmployee( db domain.IEmployee) *UpdateEmployee {
	return &UpdateEmployee{db: db}
}

func (ue *UpdateEmployee) Run (id int, employee domain.Employee) (uint, error) {
	return ue.db.Update(id, employee)
}