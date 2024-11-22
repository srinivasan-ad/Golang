package main
import(
	
)
type User struct {
	Name string `json:"name"`
	RollNo int `json:"roll no"`
	Id int `json:"id"`
}
type Product struct {
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Available bool `json:"available"`
	Name string `json:"name"`
}
type Croduct struct {
	Price float64 `json:"price"`
	Stock int `json:"stock"`
	Available bool `json:"available"`
	Name string `json:"name"`
}

type Lroduct struct {
	Available bool `json:"available"`
	Category Category `json:"category"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}
type Category struct {
	TypeName string `json:"type name"`
	Type Type `json:"type"`
}
type Type struct {
	Name string `json:"name"`
	Description string `json:"description"`
}
type Order struct {
	OrderDate string `json:"orderDate"`
	OrderID int `json:"orderID"`
	Customer Customer `json:"customer"`
	Status string `json:"status"`
}
type Customer struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Address Address `json:"address"`
}
type Address struct {
	City string `json:"city"`
	State string  `json:"state"`
	Zip string `json:"zip"`
	Street string `json:"street"`
}
func Validate(u *UserType) bool {
	return gt(len(u.Address.Street), 10) && lt(u.Address.Zip, 9999) && gt(len(u.UserName), 5) && lt(len(u.UserName), 50) && gt(u.Id, 1) && lt(u.Id, 100)
}

type UserType struct {
	UserName string `json:"userName"`
	Id int `json:"id"`
	Address UserAddressType `json:"address"`
}

type UserAddressType struct {
	Street string `json:"street"`
	Zip int `json:"zip"`
}
