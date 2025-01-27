package domain //nombre de la carpeta en la cual se encuentra el archivo

type Product struct {
	id int32
	name string
	price float32
}

//interfaces para plantilla, clases para modelos de negocio :)
//visualizar como si fuera un constructor
func NewProduct(name string, price float32) *Product { //mayúscula para public y minúscula para private
	//return &Product{id: 1, name: name, price: price}
	return &Product{name: name, price: price}
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) SetName(name string){
	p.name = name
}

func (p *Product) GetPrice() float32{
	return p.price
}

func (p *Product) SetPrice (price float32) {
	p.price = price
}