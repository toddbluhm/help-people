package models

type NeedType struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
}
