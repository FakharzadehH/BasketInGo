package handlers

import (
	"net/http"

	"github.com/FakharzadehH/BasketInGo/domain/payloads"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) AuthSignUp() echo.HandlerFunc {
	type request struct {
		payloads.SignUpRequest
	}
	type response struct {
		payloads.GenericSuccessResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res, err := h.svcs.User.SignUp(c.Request().Context(), req.SignUpRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			GenericSuccessResponse: *res,
		})
	}
}

func (h *Handlers) AuthLogin() echo.HandlerFunc {
	type request struct {
		payloads.LoginRequest
	}
	type response struct {
		payloads.LoginResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res, err := h.svcs.User.Login(c.Request().Context(), req.LoginRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(200, &response{
			LoginResponse: *res,
		})
	}
}
