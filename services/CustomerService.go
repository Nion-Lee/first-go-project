package services

import (
	"First_Go_Project/dtos"
	"First_Go_Project/repositories"
	"encoding/json"
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

func DesrilizeToStruct(str *string) (*dtos.CustomerDTO, error) {
	var dto dtos.CustomerDTO
	err := json.Unmarshal([]byte(*str), &dto)

	return &dto, err
}
