package application

import "introduccion_go/src/products/domain"



type CreateProduct struct{
	db domain.Iproduct
}

func NewCreateProduct (db domain.Iproduct) *CreateProduct{
	return &CreateProduct{db:db}
}

func (cp *CreateProduct) Run (product domain.Product) error { //se puede usar Run ó Excute para CREAR objetos de un crud
	return cp.db.Save(&product)
}