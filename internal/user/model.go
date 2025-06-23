package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex;not null;size:255"`
	Username string `json:"username" gorm:"uniqueIndex;not null;size:100"`
	Password string `json:"password" gorm:"not null"`
}
