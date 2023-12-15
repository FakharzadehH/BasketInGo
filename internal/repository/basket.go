package repository

import "gorm.io/gorm"

type basketRepository struct {
	db *gorm.DB
}

func newBasketRepository(db *gorm.DB) *basketRepository {
	return &basketRepository{db: db}
}
