package address

import "github.com/google/uuid"

type Address struct {
	ID         uuid.UUID `db:"id"`
	UserID     uuid.UUID `db:"user_id"`
	Line1      string    `db:"line1"`
	City       string    `db:"city"`
	Province   string    `db:"province"`
	PostalCode string    `db:"postal_code"`
	Country    string    `db:"country"`
	IsDefault  bool      `db:"is_default"`
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
