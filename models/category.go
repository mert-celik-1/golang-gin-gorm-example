package models

type Category struct {
	Name string `gorm:""json:"name"`
	Products []Product `json:"-"`
	ID int `json:"id" gorm:"primaryKey"`
}