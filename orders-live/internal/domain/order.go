package domain

import (
	"leal.co/orders/internal/domain/repository"
)

type Order struct {
	client          Client
	city            City
	shippingCost    int
	deliveryAddress string
	products        []Product
}

type Client struct {
	id int
}

type City struct {
	code          int
	shipphingCost int
}

type Product struct {
	code int
	name string
}

type orders struct {
	repo repository.Repository
}
