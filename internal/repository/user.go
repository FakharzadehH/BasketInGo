package repository

import (
	"github.com/FakharzadehH/BasketInGo/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetByID(id uint, user *domain.User) error {
	return u.db.First(user, id).Error
}

func (u *userRepository) GetByUsername(username string, user *domain.User) error {
	return u.db.Where("username = ?", username).First(user).Error
}

func (u *userRepository) Insert(user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) SetPassword(id uint, password string) error {
	return u.db.Model(&domain.User{}).Where("id = ?", id).Update("password", password).Error
}
