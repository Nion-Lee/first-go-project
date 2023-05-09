package services

import (
	"First_Go_Project/dtos"
	"First_Go_Project/repositories"
)

func CreateCustomer(dto *dtos.CustomerDTO) (int, error) {
	return 0, nil
}

func MapToCustomerDTO(entity *repositories.CustomerEntity) (*dtos.CustomerDTO, error) {
	return nil, nil
}

func MapToCustomerEntity(dto *dtos.CustomerDTO) (*repositories.CustomerEntity, error) {
	return nil, nil
}

func DesrilizeToStruct(json *string) *dtos.CustomerDTO {
	// var dto dtos.CustomerDTO
	// err := json.Unmarshal([]byte(*json), &dto)

	// return &dto
	return nil
}
