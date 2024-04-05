package models

type User struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	Refresh_Token string `json:"refresh_token"`
}