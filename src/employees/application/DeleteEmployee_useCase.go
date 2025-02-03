package application

import "introduccion_go/src/employees/domain"

type DeleteEmployee struct {
	db domain.IEmployee
}

func NewDeleteEmployee (db domain.IEmployee) *DeleteEmployee {
	return &DeleteEmployee{db: db}
}

func (de *DeleteEmployee) Run (id int) (uint, error) {
	return de.db.Delete(id)
}