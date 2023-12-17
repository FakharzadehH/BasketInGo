package user

import (
	"context"
	"errors"

	"github.com/FakharzadehH/BasketInGo/domain"
	"github.com/FakharzadehH/BasketInGo/domain/payloads"
	"github.com/FakharzadehH/BasketInGo/internal/helpers"
	"github.com/FakharzadehH/BasketInGo/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *User {
	return &User{repos: repos}
}

func (u *User) SignUp(ctx context.Context, payload payloads.SignUpRequest) (*payloads.GenericSuccessResponse, error) {
	err := u.repos.User.GetByUsername(payload.Username, &domain.User{})
	if err != gorm.ErrRecordNotFound {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("a user with this username already exists")
	}

	if payload.Password == "" {
		return nil, errors.New("please enter a password")
	}
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := domain.User{}
	user.Username = payload.Username
	if err := u.repos.User.Insert(&user); err != nil {
		return nil, err
	}

	if err := u.repos.User.SetPassword(user.ID, string(encryptedPass)); err != nil {
		return nil, err
	}
	return &payloads.GenericSuccessResponse{Success: true}, nil
}

func (u *User) Login(ctx context.Context, payload payloads.LoginRequest) (*payloads.LoginResponse, error) {
	user := domain.User{}
	if err := u.repos.User.GetByUsername(payload.Username, &user); err != nil {
		return nil, errors.New("Invalid Credentials")
	}
	userPassword, err := u.repos.User.GetPassword(user.ID)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(payload.Password)); err != nil {
		return nil, errors.New("Invalid Credentials")
	}
	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New("error generating jwt token")
	}
	return &payloads.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

func (u *User) IndexBaskets(ctx context.Context, userID uint) (*payloads.IndexBasketsResponse, error) {
	baskets, err := u.repos.Basket.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &payloads.IndexBasketsResponse{
		Baskets: baskets,
	}, nil
}

func (u *User) GetBasket(ctx context.Context, userID uint, basketID uint) (*payloads.GetBasketResponse, error) {
	basket := domain.Basket{}
	if err := u.repos.Basket.GetBasketByID(basketID, &basket); err != nil {
		return nil, err
	}
	if basket.UserID != userID {
		return nil, errors.New("this basket belongs to another user")
	}
	return &payloads.GetBasketResponse{
		Basket: basket,
	}, nil
}
func (u *User) UpdateBasket(ctx context.Context, userID uint, basketID uint, payload payloads.UpdateBasketRequest) (*payloads.GenericSuccessResponse, error) {
}
func (u *User) DeleteBasket(ctx context.Context, userID uint, basketID uint) (*payloads.GenericSuccessResponse, error) {
}
func (u *User) CreateBasket(ctx context.Context, payload payloads.CreateBasketRequest) (*payloads.CreateBasketResponse, error) {
}
