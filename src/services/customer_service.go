package services

import (
	"errors"
	"first-go-project/src/dtos"
	"first-go-project/src/repositories"

	"github.com/google/uuid"
)

// 測試用
func HelloWorld() string {
	return "HiHi我在這"
}

func GetCustomer(uuid *string) (*dtos.CustomerDTO, error) {
	entity := make(chan *repositories.CustomerEntity)
	err := make(chan error)

	go repositories.GetCustomerByUID(uuid, entity, err)
	dto := mapToCustomerDTO(<-entity)

	return dto, <-err
}

func CreateCustomer(dto *dtos.CustomerDTO) (string, error) {
	err := repositories.HasCustomerByUID(&dto.UID)
	if err == nil {
		return "", errors.New("已有該客戶資料，無法重新創建")
	}

	entity := mapToCustomerEntity(dto)
	uid, err := repositories.CreateCustomer(entity)

	return uid, err
}

func UpdateCustomer(dto *dtos.CustomerDTO) error {
	return nil
}

func DeleteCustomer(dto *dtos.CustomerDTO) error {
	return nil
}

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

	if entity.UID == "" {
		entity.UID = uuid.New().String()
	}

	return &entity
}
