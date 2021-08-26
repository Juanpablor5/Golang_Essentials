package usecases

import (
	"leal.co/orders/internal/domain"
	"leal.co/orders/internal/dto"
)


type Service interface {
  CreateOrder(dto.CreateOrderDTO) error
}

func NewService(domain domain.Domain) Service {
  return &service{
   domain,
  }
}

type service struct {
 orders domain.Domain
}

func (s *service) CreateOrder(input dto.CreateOrderDTO) error {
  err := s.orders.NewOrder(input)

  return err
}
