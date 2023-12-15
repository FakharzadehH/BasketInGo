package user

import "github.com/FakharzadehH/BasketInGo/internal/repository"

type User struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *User {
	return &User{repos: repos}
}
