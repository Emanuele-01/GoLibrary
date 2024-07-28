package models

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      string `json:"age"`
	Role     []Role `json:"role"`
}

type Role struct {
	IsRole string `json:"isRole"`
}
