package models

type User struct {
	Id uint32 `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}
