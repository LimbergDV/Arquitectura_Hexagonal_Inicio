package domain

//Separa el modelo de la capa de negocio
//que necesitan hacer mis actores con los casos de uso
type Iproduct interface {
	Save(product Product) (uint64, error) //podemos retornar algo acá, como un string, etc.
	GetAll() ([]Product, error) //tenemos nuestros métodos los cuales los usuarios hacen la acción con nuestra base de datos
	Delete(id int) (uint64, error)
	Update(id int, product Product) 
}