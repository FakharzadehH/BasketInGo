package user

import (
	"context"
	"errors"

	"github.com/FakharzadehH/BasketInGo/domain"
	"github.com/FakharzadehH/BasketInGo/domain/payloads"
	"github.com/FakharzadehH/BasketInGo/internal/helpers"
	"github.com/FakharzadehH/BasketInGo/internal/logger"
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
		logger.Logger().Errorw("db error while indexing baskets", "err", err)
		return nil, err
	}
	return &payloads.IndexBasketsResponse{
		Baskets: baskets,
	}, nil
}

func (u *User) GetBasket(ctx context.Context, userID uint, basketID uint) (*payloads.GetBasketResponse, error) {
	basket := domain.Basket{}
	if err := u.repos.Basket.GetBasketByID(basketID, &basket); err != nil {
		logger.Logger().Errorw("db error while getting basket", "err", err)
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

	basket := domain.Basket{}
	if err := u.repos.Basket.GetBasketByID(basketID, &basket); err != nil {
		return nil, err
	}
	if basket.UserID != userID {
		return nil, errors.New("this basket belongs to another user")
	}
	if basket.State == domain.BasketStateComplete {
		return nil, errors.New("The basket state is complete and cannot be edited anymore.")
	}
	if len(payload.Data) != 0 {
		basket.Data = payload.Data
	}
	if payload.State != "" {
		if payload.State != domain.BasketStateComplete && payload.State != domain.BasketStatePending {
			return nil, errors.New("invalid state")
		}
		basket.State = payload.State
	}
	if err := u.repos.Basket.Upsert(&basket); err != nil {
		logger.Logger().Errorw("db error while updating basket", "err", err)
		return nil, errors.New("error while updating basket")
	}
	return &payloads.GenericSuccessResponse{
		Success: true,
	}, nil
}
func (u *User) DeleteBasket(ctx context.Context, userID uint, basketID uint) (*payloads.GenericSuccessResponse, error) {
	basket := domain.Basket{}
	if err := u.repos.Basket.GetBasketByID(basketID, &basket); err != nil {
		return nil, err
	}
	if basket.UserID != userID {
		return nil, errors.New("this basket belongs to another user")
	}
	if err := u.repos.Basket.Delete(basketID); err != nil {
		logger.Logger().Errorw("db error while deleting basket", "err", err)
		return nil, errors.New("error while deleting basket")
	}
	return &payloads.GenericSuccessResponse{
		Success: true,
	}, nil
}
func (u *User) CreateBasket(ctx context.Context, userID uint, payload payloads.CreateBasketRequest) (*payloads.CreateBasketResponse, error) {
	if len(payload.Data) == 0 || (payload.State != domain.BasketStateComplete && payload.State != domain.BasketStatePending) {
		return nil, errors.New("invalid data or state")
	}
	basket := domain.Basket{
		Data:   payload.Data,
		State:  payload.State,
		UserID: userID,
	}

	if err := u.repos.Basket.Upsert(&basket); err != nil {
		logger.Logger().Errorw("db error while creating basket", "err", err)
		return nil, errors.New("error while creating basket")
	}
	return &payloads.CreateBasketResponse{
		Basket: basket,
	}, nil
}
