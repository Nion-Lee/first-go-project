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

func HasCustomerByUID(uid *string) error {
	return context.Where("UID = ?", uid).First(&CustomerEntity{}).Error
}

func CreateCustomer(entity *CustomerEntity) (string, error) {
	entity.Created = time.Now()
	entity.Updated = entity.Created

	result := context.Create(entity)
	return entity.UID, result.Error
}
