package basket

import "github.com/FakharzadehH/BasketInGo/internal/repository"

type Basket struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Basket {
	return &Basket{repos: repos}
}
