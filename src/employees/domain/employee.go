package domain

type Employee struct {
	Id int32
	First_name string
	Last_name string
	Age int32
	Curp string
	Phone_number string
	Salary float32 
}

func NewEmployee(First_name string, Last_name string, Age int32, Curp string, Phone_number string, Salary float32) *Employee{
	return &Employee{Id: 1, First_name: First_name, Last_name: Last_name, Age: Age, Curp:  Curp, Phone_number: Phone_number, Salary: Salary}
}

func (e *Employee) GetId() int32 {
	return e.Id
}

func (e *Employee) SetId(Id int32) {
	e.Id = Id
}

func (e *Employee) GetFirstName() string {
	return e.First_name
}

func (e *Employee) SetFirstName(First_name string) {
	e.First_name = First_name
} 

func (e *Employee) GetLastName() string {
	return e.Last_name
}

func (e *Employee) SetLastName (Last_name string) {
	e.Last_name = Last_name
} 

func (e *Employee) GetAge() int32 {
	return e.Age
}

func (e *Employee) SetAge (Age int32) {
	e.Age = Age
} 

func (e *Employee) GetCurp () string {
	return e.Curp
}

func (e *Employee) SetCurp (Curp string) {
	e.Curp = Curp
}

func (e *Employee) GetPhoneNumber () string {
	return e.Phone_number
}

func (e *Employee) SetPhoneNumber (Phone_number string) {
	e.Phone_number = Phone_number
}

func (e *Employee) GetSalary () float32 {
	return e.Salary
}

func (e *Employee) SetSalary (Salary float32){
	e.Salary = Salary
}

