package models

type Project struct {
	Id uint32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}
