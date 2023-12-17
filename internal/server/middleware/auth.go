package middleware

import (
	"strings"
	"time"

	"github.com/FakharzadehH/BasketInGo/domain"
	"github.com/FakharzadehH/BasketInGo/internal/helpers"
	"github.com/labstack/echo/v4"
)

func (m *Middlewares) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(header) != 2 {
				return echo.NewHTTPError(401, "Unauthorized, invalid token")
			}
			token := header[1]
			claims, err := helpers.ParseJWT(token)
			if err != nil {
				return echo.NewHTTPError(401, "Unauthorized, failed to parse token")
			}
			if claims.Expiry.Before(time.Now()) {
				return echo.NewHTTPError(401, "Unauthorized, token expired")
			}

			user := domain.User{}
			if err := m.repos.User.GetByID(claims.ID, &user); err != nil {
				return echo.NewHTTPError(401, "Unauthorized, failed to get user")
			}
			c.Set("user", &user)
			return next(c)
		}
	}
}
