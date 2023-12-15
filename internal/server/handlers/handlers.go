package handlers

import (
	"github.com/FakharzadehH/BasketInGo/internal/logger"
	"github.com/FakharzadehH/BasketInGo/internal/repository"
	"github.com/FakharzadehH/BasketInGo/service"
	"go.uber.org/zap"
)

func New(repos *repository.Repositories, svcs *service.Services) *Handlers {
	return &Handlers{
		repos: repos,
		svcs:  svcs,
		log:   logger.Logger(),
	}
}

type Handlers struct {
	repos *repository.Repositories
	svcs  *service.Services
	log   *zap.SugaredLogger
}
