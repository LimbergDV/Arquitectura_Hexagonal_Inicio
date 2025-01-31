package application

import "introduccion_go/src/products/domain"

type GetAllProduct struct {
	db domain.Iproduct
}

func NewGetAllProduct(db domain.Iproduct) *GetAllProduct  {
	return &GetAllProduct{db: db}
}

func (lp *GetAllProduct) Run() []domain.Product{
	return lp.db.GetAll()
}


