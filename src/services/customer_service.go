package services

import (
	"errors"
	"first-go-project/src/dtos"
	"first-go-project/src/repositories"
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

func CreateCustomer(dto *dtos.CustomerDTO) (*string, error) {
	err := make(chan error)
	go repositories.HasCustomerByUID(&dto.UID, err)

	if <-err == nil {
		return nil, errors.New("已有該客戶資料，無法重新創建")
	}

	uid := make(chan *string)
	entity := mapToCustomerEntity(dto)
	go repositories.CreateCustomer(entity, uid, err)

	return <-uid, <-err
}

func UpdateCustomer(dto *dtos.CustomerDTO) error {
	return nil
}

func DeleteCustomer(dto *dtos.CustomerDTO) error {
	return nil
}
