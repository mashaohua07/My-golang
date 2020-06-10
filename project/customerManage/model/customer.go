package model

import "fmt"

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer  {
	return Customer{
		id,
		name,
		gender,
		age,
		phone,
		email,
	}
}

//返回用户的信息
func (this Customer) GetInfo()  string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}