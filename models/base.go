package models

import "github.com/gocraft/dbr"

type BaseModel struct {
	Id        int64        `json:"id"`
	CreatedAt dbr.NullTime `json:"created_at"`
	UpdatedAt dbr.NullTime `json:"updated_at"`
	DeletedAt dbr.NullTime `json:"deleted_at"`
}
