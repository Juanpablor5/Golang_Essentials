package repository

type Repository interface {
  SaveOrder(NewOrderDTO) error
  GetCity(int) CityDTO
  GetClient(int) ClientDTO
}
