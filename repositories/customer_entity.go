package repositories

import "time"

type CustomerEntity struct {
	ID      int    `gorm:"primaryKey"`
	UID     string `gorm:"type:uuid;not null;"`
	Name    string
	Age     int
	Email   string
	Created time.Time
	Updated time.Time
}
