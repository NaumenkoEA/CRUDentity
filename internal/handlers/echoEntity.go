package handlers

import (
	"github.com/EgMeln/CRUDentity/internal/repository/postgreSQL"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Advert struct {
	ID      int     `json:"id" form:"id" query:"id"`
	Address string  `json:"address" form:"address" query:"address"`
	Price   float32 `json:"price" form:"price" query:"price"`
}

func Create(e echo.Context) (err error) {
	c := new(Advert)
	if err = e.Bind(c); err != nil {
		return err
	}
	advert := Advert{
		ID:      c.ID,
		Address: c.Address,
		Price:   c.Price,
	}
	postgreSQL.CreateRecord(advert.ID, advert.Address, advert.Price)
	return e.JSON(http.StatusOK, c)
}

func ReadAll(e echo.Context) error {
	return e.String(http.StatusOK, postgreSQL.ReadAllRecords())
}
func ReadById(e echo.Context) (err error) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return err
	}
	return e.String(http.StatusOK, postgreSQL.ReadRecordByNum(id))
}
func UpdateRecord(e echo.Context) (err error) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return err
	}
	c := new(Advert)
	if err = e.Bind(c); err != nil {
		return err
	}
	advert := Advert{
		Address: c.Address,
		Price:   c.Price,
	}
	postgreSQL.UpdateRecord(id, advert.Address, advert.Price)
	return e.JSON(http.StatusOK, e)
}

func Delete(e echo.Context) (err error) {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return err
	}
	postgreSQL.DeleteRecord(id)
	return e.JSON(http.StatusOK, e)
}
