package application

import "introduccion_go/src/products/domain"

type UpdateProduct struct {
	db domain.Iproduct
}

func NewUpdateProduct (db domain.Iproduct) *UpdateProduct {
	return &UpdateProduct{db:db}
}

func (up *UpdateProduct) Run (id int, product domain.Product) (uint, error){
	return up.db.Update(id, product)
}