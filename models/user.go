package models

type User struct {
	ID       int    `json:"-" gorm:"id"`
	Name     string `json:"name" gorm:"name" binding:"required"`
	Username string `json:"username" gorm:"username" binding:"required"`
	Password string `json:"password" gorm:"password_hash" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"username" binding:"required"`
	Password string `json:"password" gorm:"password" binding:"required"`
}
