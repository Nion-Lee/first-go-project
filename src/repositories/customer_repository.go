package repositories

import "time"

func GetCustomerByUID(uid *string) (*CustomerEntity, error) {
	created := &CustomerEntity{}
	result := context.Where("UID = ?", uid).First(created)

	return created, result.Error
}

func HasCustomerByUID(uid *string) error {
	aa := context.Where("UID = ?", uid).First(&CustomerEntity{}).Error
	return aa
}

func CreateCustomer(entity *CustomerEntity) (string, error) {
	entity.Created = time.Now()
	entity.Updated = entity.Created

	result := context.Create(entity)
	return entity.UID, result.Error
}
