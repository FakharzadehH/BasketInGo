package middleware

import (
	"github.com/FakharzadehH/BasketInGo/internal/logger"
	"github.com/FakharzadehH/BasketInGo/internal/repository"
	"github.com/FakharzadehH/BasketInGo/service"
	"go.uber.org/zap"
)

func NewMiddlewares(svcs *service.Services, repos *repository.Repositories) *Middlewares {
	return &Middlewares{
		svcs:  svcs,
		repos: repos,
		log:   logger.Logger(),
	}
}

type Middlewares struct {
	svcs  *service.Services
	repos *repository.Repositories
	log   *zap.SugaredLogger
}
