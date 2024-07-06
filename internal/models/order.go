package models

import (
	"time"
)

type OrderId int64
type AddresseeId int64

type Order struct {
	OrderId     OrderId     `json:"orderId" db:"orderid"`         // ID заказа
	AddresseeId AddresseeId `json:"addresseeId" db:"addresseeid"` // ID получателя
	ShelfLife   time.Time   `json:"shelfLife" db:"shelflife"`     // Срок хранения заказа
	Weight      int
	Price       int
	Wrapper     Wrapper
}
