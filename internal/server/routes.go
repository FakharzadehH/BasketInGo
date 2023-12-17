package server

import (
	"github.com/FakharzadehH/BasketInGo/internal/server/handlers"
	"github.com/FakharzadehH/BasketInGo/internal/server/middleware"
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo, h *handlers.Handlers, m *middleware.Middlewares) {
	user := e.Group("/user", m.Auth())
	user.GET("/basket", h.UserIndexBaskets())
	user.GET("/basket/:id", h.UserGetBasket())
	user.POST("/basket", h.UserCreateBasket())
	user.PATCH("/basket/:id", h.UserUpdateBasket())
	user.DELETE("/basket/:id", h.UserDeleteBasket())

	auth := e.Group("/auth")
	auth.POST("/signup", h.AuthSignUp())
	auth.POST("/login", h.AuthLogin())

}
