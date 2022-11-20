package handlers

import (
	"github.com/EgMeln/CRUDentity/internal/model"
	"github.com/EgMeln/CRUDentity/internal/request"
	"github.com/EgMeln/CRUDentity/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AdvertHandler struct {
	service *service.AdvertService
}

func NewServiceAdvert(srv *service.AdvertService) AdvertHandler {
	return AdvertHandler{service: srv}
}

func (handler *AdvertHandler) Add(e echo.Context) (err error) {
	c := new(request.AdvertCreate)
	if err = e.Bind(c); err != nil {
		return e.JSON(http.StatusBadRequest, c)
	}
	err = handler.service.Add(e.Request().Context(), &model.Advert{ID: c.ID, Address: c.Address, Price: c.Price})
	if err != nil {
		return e.JSON(http.StatusBadRequest, c)
	}
	return e.JSON(http.StatusOK, c)
}

func (handler *AdvertHandler) GetAll(e echo.Context) error {
	parkingLots, err := handler.service.GetAll(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusBadRequest, parkingLots)
	}
	return e.JSON(http.StatusOK, parkingLots)
}

func (handler *AdvertHandler) GetByID(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return err
	}
	var advert *model.Advert
	advert, err = handler.service.GetByID(e.Request().Context(), id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, advert)
	}
	return e.JSON(http.StatusOK, advert)
}

func (handler *AdvertHandler) Update(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return err
	}
	c := new(request.AdvertUpdate)
	if err = e.Bind(c); err != nil {
		return err
	}
	err = handler.service.Update(e.Request().Context(), id, c.Address, c.Price)
	if err != nil {
		return e.JSON(http.StatusBadRequest, c)
	}
	return e.JSON(http.StatusOK, c)
}

func (handler *AdvertHandler) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return err
	}
	err = handler.service.Delete(e.Request().Context(), id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, e)
	}
	return e.JSON(http.StatusOK, e)
}
