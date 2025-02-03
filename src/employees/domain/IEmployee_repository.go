package domain

//Separa el modelo de la capa de negocio
//que necesitan hacer mis actores con los casos de uso
type IEmployee interface {
	Save(employee Employee) (uint, error) //podemos retornar algo acá, como un string, etc.
	GetAll() ([]Employee) //tenemos nuestros métodos los cuales los usuarios hacen la acción con nuestra base de datos
	Delete(id int) (uint, error)
	Update(id int, employee Employee) (uint, error)
}