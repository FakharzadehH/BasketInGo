package payloads

import (
	"time"

	"github.com/FakharzadehH/BasketInGo/domain"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string      `json:"token"`
	ExpiresAt *time.Time  `json:"expires_at"`
	User      domain.User `json:"user"`
}

type CreateBasketRequest struct {
	Data  string      `json:"data"`
	State basketState `json:"state"`
}

type UpdateBasketRequest struct {
	Data  string      `json:"data"`
	State basketState `json:"state"`
}

type GetBasketResponse struct {
	domain.Basket `json:"basket"`
}

type IndexBasketsResponse struct {
	Baskets []domain.Basket `json:"baskets"`
}

type basketState string

const (
	basketStatePending  basketState = "PENDING"
	basketStateComplete basketState = "COMPLETED"
)
