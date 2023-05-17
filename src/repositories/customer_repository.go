package repositories

import (
	"time"
)

func GetCustomerByUID(uid *string, entity chan *CustomerEntity, err chan error) {
	created := &CustomerEntity{}
	result := context.Where("UID = ?", uid).First(created)

	entity <- created
	err <- result.Error
}

func HasCustomerByUID(uid *string, err chan error) {
	err <- context.Where("UID = ?", uid).First(&CustomerEntity{}).Error
}

func CreateCustomer(entity *CustomerEntity, uid chan *string, err chan error) {
	entity.Created = time.Now()
	entity.Updated = entity.Created

	result := context.Create(entity)

	uid <- &entity.UID
	err <- result.Error
}

func UpdateCustomer(entity *CustomerEntity, err chan error) {
	result := context.Model(&CustomerEntity{}).Where("uid = ?", entity.UID).Updates(CustomerEntity{
		Name:    entity.Name,
		Age:     entity.Age,
		Email:   entity.Email,
		Updated: time.Now(),
	})

	err <- result.Error
}
