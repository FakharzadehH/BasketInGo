package payloads

import (
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
	Token string      `json:"token"`
	User  domain.User `json:"user"`
}

type CreateBasketRequest struct {
	Data  map[string]interface{} `json:"data"`
	State domain.BasketState     `json:"state"`
}
type CreateBasketResponse struct {
	domain.Basket `json:"basket"`
}

type UpdateBasketRequest struct {
	Data  map[string]interface{} `json:"data"`
	State domain.BasketState     `json:"state"`
}

type GetBasketResponse struct {
	domain.Basket `json:"basket"`
}

type IndexBasketsResponse struct {
	Baskets []domain.Basket `json:"baskets"`
}
