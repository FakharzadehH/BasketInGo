package server

import (
	"github.com/FakharzadehH/BasketInGo/internal/server/handlers"
	"github.com/FakharzadehH/BasketInGo/internal/server/middleware"
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo, handler *handlers.Handlers, m *middleware.Middlewares) {
	// := e.Group("/basket")

	// basket.GET("")
	// basket.POST("")
	// basket.PATCH("")
	// basket.DELETE("")
}
