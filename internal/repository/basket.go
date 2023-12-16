package repository

import (
	"github.com/FakharzadehH/BasketInGo/domain"
	"gorm.io/gorm"
)

type basketRepository struct {
	db *gorm.DB
}

func newBasketRepository(db *gorm.DB) *basketRepository {
	return &basketRepository{db: db}
}

func (b *basketRepository) GetBasketByID(id uint, basket *domain.Basket) error {
	return b.db.First(basket, id).Error
}

func (b *basketRepository) GetByUserID(user_id uint) ([]domain.Basket, error) {
	baskets := []domain.Basket{}
	return baskets, b.db.Where("user_id = ?", user_id).Find(&baskets).Error
}

func (b *basketRepository) Upsert(basket *domain.Basket) error {
	return b.db.Save(basket).Error
}

func (b *basketRepository) Delete(id uint) error {
	return b.db.Delete(&domain.Basket{}, id).Error
}
