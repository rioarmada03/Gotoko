package app

import "github.com/RioArmada3989/gotoko/app/controllers/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
		{Model: models.Order{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderCustomer{}},
		{Model: models.Payment{}},
		{Model: models.Shipment{}},
	}
}
