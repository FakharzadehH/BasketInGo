package repository

import (
	"github.com/FakharzadehH/BasketInGo/domain"
	"gorm.io/gorm"
)

type Repositories struct {
	User   User
	Basket Basket
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:   newUserRepository(db),
		Basket: newBasketRepository(db),
	}
}

type User interface {
	GetByID(id uint, user *domain.User) error
	GetByUsername(username string, user *domain.User) error
	Insert(user *domain.User) error
	SetPassword(id uint, password string) error
	GetPassword(id uint) (string, error)
}

type Basket interface {
	GetBasketByID(id uint, basket *domain.Basket) error
	GetByUserID(user_id uint) ([]domain.Basket, error)
	Upsert(basket *domain.Basket) error
	Delete(id uint) error
}
