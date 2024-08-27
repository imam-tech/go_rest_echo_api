package services

import (
	"fmt"
	"rest_echo_api/domain/item/models"
	"rest_echo_api/domain/item/repositories"
	"rest_echo_api/helpers"

	"gorm.io/gorm"
)

type itemService struct {
	itemRepo repositories.ItemRepository
}

func (service *itemService) Create(item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Create(item); err != nil {
		response.Status = 500
		response.Messages = "Failed to create a new item"
	} else {
		response.Status = 200
		response.Messages = "Success to create a new item"
	}
	return response
}

func (service *itemService) Update(idItem int, item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Update(idItem, item); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to update item: ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to update item "
	}
	return response
}

func (service *itemService) Delete(idItem int) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.Delete(idItem); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to delete item", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to delete item"
	}
	return response
}

func (service *itemService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.GetAll()
	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}

	return response
}

func (service *itemService) GetById(idItem int) helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.GetById(idItem)
	if err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get item: ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to get item"
		response.Data = data
	}

	return response
}

type ItemService interface {
	Create(item models.Item) helpers.Response
	Update(idItem int, item models.Item) helpers.Response
	Delete(idItem int) helpers.Response
	GetById(idItem int) helpers.Response
	GetAll() helpers.Response
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService{itemRepo: repositories.NewItemRepository(db)}
}
