package repositories

import "github.com/google/uuid"

type CustomerEntity struct {
	ID    int
	UID   uuid.UUID
	Name  string
	Age   int
	Email string
}
