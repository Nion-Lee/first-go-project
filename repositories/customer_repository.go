package repositories

func GetCustomerByUID(uid *string) (*CustomerEntity, error) {
	return nil, nil
}

func CreateCustomer(entity *CustomerEntity) (*CustomerEntity, error) {

	// 等待orm 套件安裝
	return &CustomerEntity{
		ID:    1,
		UID:   entity.UID,
		Name:  entity.Name,
		Age:   entity.Age,
		Email: entity.Email,
	}, nil
}
