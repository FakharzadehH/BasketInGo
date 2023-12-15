package repository

import "gorm.io/gorm"

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
}

type Basket interface{}
