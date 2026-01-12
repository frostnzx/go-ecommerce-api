package address

import "github.com/google/uuid"

type Address struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Line1      string
	City       string
	Province   string
	PostalCode string
	Country    string
	IsDefault  bool
}

func New(userId uuid.UUID, line1, city, province, postalCode, country string) Address {
	return Address{
		ID:         uuid.New(),
		UserID:     userId,
		Line1:      line1,
		City:       city,
		Province:   province,
		PostalCode: postalCode,
		Country:    country,
		IsDefault:  false,
	}
}
