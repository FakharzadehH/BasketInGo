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
func (u *userRepository) GetPassword(id uint) (string, error) {
	user := struct {
		Password string `gorm:"column:password"`
		ID       uint   `gorm:"column:id"`
	}{}
	err := u.db.Model(&domain.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}
