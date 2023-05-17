package services

import (
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

func CreateCustomer(dto *dtos.CustomerDTO) (*string, error) {
	err := make(chan error)
	uid := make(chan *string)

	entity := mapToCustomerEntity(dto)
	entity.UID = uuid.New().String()
	go repositories.CreateCustomer(entity, uid, err)

	return <-uid, <-err
}

func UpdateCustomer(dto *dtos.CustomerDTO) error {
	err := make(chan error)
	entity := mapToCustomerEntity(dto)

	go repositories.UpdateCustomer(entity, err)

	return <-err
}

func DeleteCustomer(uuid *string) error {
	err := make(chan error)
	go repositories.DeleteCustomer(uuid, err)

	return <-err
}
