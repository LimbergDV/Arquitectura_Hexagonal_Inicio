package domain

//Separa el modelo de la capa de negocio
//que necesitan hacer mis actores con los casos de uso
type Iproduct interface {
	Save(product *Product) error //podemos retornar algo acá, como un string, etc.
	GetAll([] Product, error) //tenemos nuestros métodos los cuales los usuarios hacen la acción con nuestra base de datos
}