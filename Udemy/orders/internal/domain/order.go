package domain

import (
	"fmt"

	"leal.co/orders/internal/domain/repository"
	"leal.co/orders/internal/dto"
)

type Order struct {
	client          Client
	city            City
	shippingCost    int
	deliveryAddress string
	products        []Product
}

type Client struct {
	Id int
}

type City struct {
	Code          int
	ShipphingCost int
}

type Product struct {
	code int
	name string
}

type orders struct {
	repo repository.Repository
}

func NewDomain(repo repository.Repository) Domain {
	return &orders{
		repo: repo,
	}
}

func (o *orders) NewOrder(input dto.CreateOrderDTO) error {

	city := o.repo.GetCity(input.CityCode)
	if city.Code == 0 {
		return fmt.Errorf("city does not exist")
	}

	client := o.repo.GetClient(input.ClientId)
	if client.Id == 0 {
		return fmt.Errorf("client does not exist")
	}

	productsIds := make([]int, len(input.Products))
	for i, product := range input.Products {
		productsIds[i] = product.Code
	}

	newOrderRepoDTO := repository.NewOrderDTO{
		ClientId:        client.Id,
		CityCode:        input.CityCode,
		DeliveryAddress: input.DeliveryAddress,
		Products:        productsIds,
		DeliveryCost:    city.DeliveryCost,
	}

	err := o.repo.SaveOrder(newOrderRepoDTO)

	return err
}
