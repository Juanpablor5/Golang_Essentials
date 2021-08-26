package domain

import "leal.co/orders/internal/dto"

type Domain interface {
  NewOrder(dto.CreateOrderDTO) error
}
