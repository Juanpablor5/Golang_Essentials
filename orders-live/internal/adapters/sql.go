package adapters

import (
	"database/sql"
	"fmt"

	"leal.co/orders/internal/domain/repository"
)

type sqlRepository struct {
  db *sql.DB
}

func NewSqlRepository(db *sql.DB) repository.Repository {
  return &sqlRepository{
    db,
  }
}

func (s *sqlRepository) SaveOrder(orders repository.NewOrderDTO) error {
  var id int
	if err := s.db.QueryRow("INSERT INTO orders (client_id, city_code, delivery_address, delivery_cost) VALUES ($1, $2, $3, $4) returning id",
		orders.ClientId, orders.CityCode, orders.DeliveryAddress, orders.DeliveryCost).
		Scan(&id); err != nil {
		return fmt.Errorf("error creating user, %v", err)
	}

  for _, product := range orders.Products {
    _, err := s.db.Exec("INSERT INTO orders_products (order_id, product_id) VALUES ($1, $2)",
    id, product)
    if err != nil {
      return fmt.Errorf("error saving product: %d", product)
    }
  }

  return nil
}

func (s *sqlRepository) GetCity(code int) repository.CityDTO {
  var ans repository.CityDTO
  err := s.db.QueryRow("SELECT * FROM cities WHERE code = $1", code).
  Scan(&ans.Code, &ans.DeliveryCost)
  if err != nil {
    fmt.Printf("could not get cities: %v", err)
    return ans
  }

  return ans
}

func (s *sqlRepository) GetClient(id int) repository.ClientDTO {
  var ans repository.ClientDTO

  err := s.db.QueryRow("SELECT * FROM clients WHERE id = $1", id).
  Scan(&ans.Id)
  if err != nil {
    fmt.Printf("could not get client: %v", err)
    return ans
  }

  return ans

}


