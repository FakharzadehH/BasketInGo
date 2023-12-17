package handlers

import (
	"net/http"
	"strconv"

	"github.com/FakharzadehH/BasketInGo/domain"
	"github.com/FakharzadehH/BasketInGo/domain/payloads"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) UserIndexBaskets() echo.HandlerFunc {
	type response struct {
		payloads.IndexBasketsResponse
	}
	return func(c echo.Context) error {
		user := c.Get("user").(*domain.User)
		res, err := h.svcs.User.IndexBaskets(c.Request().Context(), user.ID)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			IndexBasketsResponse: *res,
		})
	}
}

func (h *Handlers) UserGetBasket() echo.HandlerFunc {
	type response struct {
		payloads.GetBasketResponse
	}
	return func(c echo.Context) error {
		user := c.Get("user").(*domain.User)
		basketID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id format")
		}

		res, err := h.svcs.User.GetBasket(c.Request().Context(), user.ID, uint(basketID))
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			GetBasketResponse: *res,
		})
	}
}

func (h *Handlers) UserUpdateBasket() echo.HandlerFunc {
	type request struct {
		payloads.UpdateBasketRequest
	}
	type response struct {
		payloads.GenericSuccessResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, "invalid request")
		}
		user := c.Get("user").(*domain.User)
		basketID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id format")
		}
		res, err := h.svcs.User.UpdateBasket(c.Request().Context(), user.ID, uint(basketID), req.UpdateBasketRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			GenericSuccessResponse: *res,
		})
	}
}

func (h *Handlers) UserDeleteBasket() echo.HandlerFunc {
	type response struct {
		payloads.GenericSuccessResponse
	}
	return func(c echo.Context) error {
		//TODO : extract user id from JWT
		user := c.Get("user").(*domain.User)
		basketID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid id format")
		}
		res, err := h.svcs.User.DeleteBasket(c.Request().Context(), user.ID, uint(basketID))
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			GenericSuccessResponse: *res,
		})
	}
}

func (h *Handlers) UserCreateBasket() echo.HandlerFunc {
	type request struct {
		payloads.CreateBasketRequest
	}
	type response struct {
		payloads.CreateBasketResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, "bad request format")
		}
		user := c.Get("user").(*domain.User)
		res, err := h.svcs.User.CreateBasket(c.Request().Context(), user.ID, req.CreateBasketRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			CreateBasketResponse: *res,
		})
	}
}
