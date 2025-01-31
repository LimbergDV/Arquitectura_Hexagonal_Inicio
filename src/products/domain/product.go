package domain //nombre de la carpeta en la cual se encuentra el archivo

type Product struct {
	Id int32 `json:"id,omitempty"` 
	Name string `json:"name"`
	Price float32 `json:"price"`
}

//interfaces para plantilla, clases para modelos de negocio :)
//visualizar como si fuera un constructor
func NewProduct(Name string, Price float32) *Product { //mayúscula para public y minúscula para private
	//return &Product{id: 1, name: name, price: price}
	return &Product{Id: 1, Name: Name, Price: Price}
}

//getters and setters 
func (p *Product) GetId() int32 {
	return p.Id
}

func (p *Product) SetId(Id int32) {
	p.Id = Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string){
	p.Name = name
}

func (p *Product) GetPrice() float32{
	return p.Price
}

func (p *Product) SetPrice (price float32) {
	p.Price = price
}