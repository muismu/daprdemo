package common

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id              uuid.UUID `json:"Id"`
	Amount          int64     `json:"Amount"`
	Customer        string    `json:"Customer"`
	Date            time.Time `json:"Date"`
	ShippingAddress string    `json:"ShippingAddress"`
	Category        string    `json:"Category"`
	Provider        string    `json:"Provider"`
}
