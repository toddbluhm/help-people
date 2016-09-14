package models

import "github.com/gocraft/dbr"

type Need struct {
	BaseModel
	Name        dbr.NullString `json:"name"`
	NeedType    string         `json:"type" gorm:"not null"`
	Description dbr.NullString `json:"description"`
}
