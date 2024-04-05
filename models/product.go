package models

type Product struct {
	ID         int      `json:"id" gorm:"primaryKey"`
	Name       string   `gorm:""json:"name"`
	Quantity   int      `json:"quantity"`
	IsDeleted  bool     `json:"isDeleted"`
	CategoryId uint     `json:"categoryId"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}