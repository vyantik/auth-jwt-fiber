package user

import (
	"github.com/vyantik/auth-jwt-service/pkg/db"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(user User) (*User, error) {
	if err := r.db.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	var user User
	result := r.db.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) GetByID(id uint) (*User, error) {
	var user User
	result := r.db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}