package application

import "introduccion_go/src/products/domain"

type DeleteProduct struct{
	db domain.Iproduct
}

func NewDeleteProduct(db domain.Iproduct) *DeleteProduct{
	return &DeleteProduct{db: db}
}

func (dp *DeleteProduct ) Run(id int) error{
	return dp.db.Delete(id)
}