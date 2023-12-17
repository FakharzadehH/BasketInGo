package service

import (
	"context"

	"github.com/FakharzadehH/BasketInGo/domain/payloads"
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
	SignUp(ctx context.Context, payload payloads.SignUpRequest) (*payloads.GenericSuccessResponse, error)
	Login(ctx context.Context, payload payloads.LoginRequest) (*payloads.LoginResponse, error)
	IndexBaskets(ctx context.Context, userID uint) (*payloads.IndexBasketsResponse, error)
	GetBasket(ctx context.Context, userID uint, basketID uint) (*payloads.GetBasketResponse, error)
	UpdateBasket(ctx context.Context, userID uint, basketID uint, payload payloads.UpdateBasketRequest) (*payloads.GenericSuccessResponse, error)
	DeleteBasket(ctx context.Context, userID uint, basketID uint) (*payloads.GenericSuccessResponse, error)
	CreateBasket(ctx context.Context, userID uint, payload payloads.CreateBasketRequest) (*payloads.CreateBasketResponse, error)
}
