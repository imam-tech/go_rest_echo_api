package controllers

import (
	"net/http"
	"rest_echo_api/domain/item/models"
	"rest_echo_api/domain/item/services"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ItemController struct {
	itemSerice services.ItemService
}

func (controller ItemController) Create(c echo.Context) error {
	type payload struct {
		NamaItem    string  `json:"nama_item" validate:"required"`
		Unit        string  `json:"unit" validate:"required"`
		Stok        int     `json:"stok" validate:"required"`
		HargaSatuan float64 `json:"harga_satuan" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	result := controller.itemSerice.Create(
		models.Item{
			NamaItem:    payloadValidator.NamaItem,
			Unit:        payloadValidator.Unit,
			Stok:        payloadValidator.Stok,
			HargaSatuan: payloadValidator.HargaSatuan,
		},
	)

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) Update(c echo.Context) error {
	type payload struct {
		NamaItem    string  `json:"nama_item" validate:"required"`
		Unit        string  `json:"unit" validate:"required"`
		Stok        int     `json:"stok" validate:"required"`
		HargaSatuan float64 `json:"harga_satuan" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	idItem, _ := strconv.Atoi(c.Param("id_item"))
	result := controller.itemSerice.Update(
		idItem,
		models.Item{
			NamaItem:    payloadValidator.NamaItem,
			Unit:        payloadValidator.Unit,
			Stok:        payloadValidator.Stok,
			HargaSatuan: payloadValidator.HargaSatuan,
		},
	)

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) Delete(c echo.Context) error {
	idItem, _ := strconv.Atoi(c.Param("id_item"))
	result := controller.itemSerice.Delete(idItem)

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) GetAll(c echo.Context) error {
	result := controller.itemSerice.GetAll()

	return c.JSON(http.StatusOK, result)
}

func (controller ItemController) GetById(c echo.Context) error {
	idItem, _ := strconv.Atoi(c.Param("id_item"))
	result := controller.itemSerice.GetById(idItem)

	return c.JSON(http.StatusOK, result)
}

func NewItemController(db *gorm.DB) ItemController {
	service := services.NewItemService(db)
	controller := ItemController{
		itemSerice: service,
	}

	return controller
}
