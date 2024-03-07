package models

import "time"

type Bug struct {
	Id uint32 `json:"id"`
	Description string `json:"description"`
	CreationDate time.Time `json:"creation_date"`
	UserId uint32 `json:"user_id"`
	ProjectId uint32 `json:"project_id"`
}
