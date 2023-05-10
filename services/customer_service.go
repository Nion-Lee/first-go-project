package services

import (
	"First_Go_Project/dtos"
	"First_Go_Project/repositories"
	"encoding/json"
)

func HelloWorld() string {
	return "HiHi我在這"
}

func GetCustomer(dto *dtos.CustomerDTO) (*dtos.CustomerDTO, error) {
	return nil, nil
}

func CreateCustomer(dto *dtos.CustomerDTO) (*dtos.CustomerDTO, error) {
	entity, err := getCustomerByUID(&dto.UID)
	if entity != nil || err != nil {
		return nil, err
	}

	newEntity := MapToCustomerEntity(dto)
	newEntity, err = repositories.CreateCustomer(newEntity)
	newDTO := MapToCustomerDTO(newEntity)

	return newDTO, err
}

func UpdateCustomer(dto *dtos.CustomerDTO) error {
	return nil
}

func DeleteCustomer(dto *dtos.CustomerDTO) error {
	return nil
}

func MapToCustomerDTO(entity *repositories.CustomerEntity) *dtos.CustomerDTO {
	dto := dtos.CustomerDTO{
		UID:   entity.UID,
		Name:  entity.Name,
		Age:   entity.Age,
		Email: entity.Email,
	}

	return &dto
}

func MapToCustomerEntity(dto *dtos.CustomerDTO) *repositories.CustomerEntity {
	entity := repositories.CustomerEntity{
		UID:   dto.UID,
		Name:  dto.Name,
		Age:   dto.Age,
		Email: dto.Email,
	}

	return &entity
}

func DesrilizeToStruct(str *string) (*dtos.CustomerDTO, error) {
	var dto dtos.CustomerDTO
	err := json.Unmarshal([]byte(*str), &dto)

	return &dto, err
}

func getCustomerByUID(uid *string) (*repositories.CustomerEntity, error) {
	return repositories.GetCustomerByUID(uid)
}
