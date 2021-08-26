package domain

import (
	"testing"

	"leal.co/orders/internal/domain/repository"
	"leal.co/orders/internal/dto"
)

func TestNewOrder(t *testing.T) {
	_ = dto.CreateOrderDTO{
		ClientId:        123,
		CityCode:        123456,
		DeliveryAddress: "address",
		Products: []dto.ProductDTO{
			{123, "prod 1"},
			{456, "prod 2"},
		},
	}

	_ = repository.NewOrderDTO{
		ClientId:        123,
		CityCode:        123456,
		DeliveryAddress: "address",
		Products:        []int{123, 456},
		DeliveryCost:    789,
	}

	t.Run("creates the order correctly", func(t *testing.T) {
	})

	t.Run("should return error if the client doesn't exist", func(t *testing.T) {
	})

	t.Run("should return error if the city doesn't exist", func(t *testing.T) {
	})
}
