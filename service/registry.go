package service

import (
	"github.com/FakharzadehH/BasketInGo/internal/repository"
	"github.com/FakharzadehH/BasketInGo/service/basket"
	"github.com/FakharzadehH/BasketInGo/service/user"
)

type Services struct {
	User   User
	Basket Basket
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User:   user.New(repos),
		Basket: basket.New(repos),
	}
}

type User interface {
}
type Basket interface {
}
