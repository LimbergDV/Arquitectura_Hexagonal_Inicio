package validators

import (
	"errors"
	"introduccion_go/src/products/domain"
)

func CheckProduct(product domain.Product) error{
	if product.Name == "" {
		return errors.New("No puedes poner un nombre vacio")
	}

	if product.Price < 0{
		return errors.New("El precio debe ser mayor a 0")
	}

	if product.Id < 0{
		return errors.New("El id debe de ser mayor a 0")
	}
	
	return nil
}