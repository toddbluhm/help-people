package model

import "github.com/gocraft/dbr"

type BaseModel struct {
	Id        int64
	CreatedAt dbr.NullTime
	UpdatedAt dbr.NullTime
	DeletedAt dbr.NullTime
}
