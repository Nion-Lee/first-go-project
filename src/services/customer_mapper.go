package services

import (
	"first-go-project/src/dtos"
	"first-go-project/src/repositories"
)

func mapToCustomerDTO(entity *repositories.CustomerEntity) *dtos.CustomerDTO {
	dto := dtos.CustomerDTO{
		UID:   entity.UID,
		Name:  entity.Name,
		Age:   entity.Age,
		Email: entity.Email,
	}

	return &dto
}

func mapToCustomerEntity(dto *dtos.CustomerDTO) *repositories.CustomerEntity {
	entity := repositories.CustomerEntity{
		UID:   dto.UID,
		Name:  dto.Name,
		Age:   dto.Age,
		Email: dto.Email,
	}

	return &entity
}
