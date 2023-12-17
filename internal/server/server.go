package server

import (
	"github.com/FakharzadehH/BasketInGo/internal/config"
	"github.com/FakharzadehH/BasketInGo/internal/repository"
	"github.com/FakharzadehH/BasketInGo/internal/server/handlers"
	"github.com/FakharzadehH/BasketInGo/internal/server/middleware"
	"github.com/FakharzadehH/BasketInGo/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Start() error {
	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()
	db, err := config.NewGORMConnection(config.GetConfig())
	if err != nil {
		return err
	}

	repos := repository.NewRepositories(db)
	svcs := service.NewServices(repos)
	middlewares := middleware.NewMiddlewares(svcs, repos)
	handler := handlers.New(repos, svcs)
	routes(e, handler, middlewares)
	return e.Start(":1323")
}
