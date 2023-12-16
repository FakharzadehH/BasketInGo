package service

import (
	"github.com/FakharzadehH/BasketInGo/internal/repository"
	"github.com/FakharzadehH/BasketInGo/service/user"
)

type Services struct {
	User User
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User: user.New(repos),
	}
}

type User interface {
}
