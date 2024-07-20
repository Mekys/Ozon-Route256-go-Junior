package storage

import (
	"errors"
	"homework-3/internal/hash"
	"homework-3/internal/models"
	"time"
)

type OrderStatus int8

const (
	ReturnToDeliverer   = OrderStatus(iota - 2) // Заказ возвращен курьеру
	ReturnFromAddressee                         // Принят возврат то клиента
	InStock                                     // На складе
	GivenToAddressee                            // Выдан клиенту
)
const (
	ReturnFromAddresseeExpiration = 1 // Принят возврат то клиента
)

type Order struct {
	OrderId       models.OrderId     `json:"orderId" db:"orderid"`         // ID заказа
	AddresseeId   models.AddresseeId `json:"addresseeId" db:"addresseeid"` // ID получателя
	ShelfLife     time.Time          `json:"shelfLife" db:"shelflife"`     // Срок хранения заказа
	Status        OrderStatus        `json:"order_status" db:"order_status"`
	StatusUpdated time.Time          `json:"status_updated_date" db:"status_updated_date"` // Последнее обновление статуса
	Hash          string             `json:"hash_code" db:"hash_code"`
	Weight        int                `json:"weight" db:"weight"`
	Price         int                `json:"price" db:"price"`
}

func (t Order) toDomain() models.Order {
	return models.Order{
		OrderId:     t.OrderId,
		AddresseeId: t.AddresseeId,
		ShelfLife:   t.ShelfLife,
		Weight:      t.Weight,
		Price:       t.Price,
	}
}

func transform(order models.Order) Order {
	return Order{
		Status:        InStock,
		OrderId:       order.OrderId,
		ShelfLife:     order.ShelfLife,
		AddresseeId:   order.AddresseeId,
		Weight:        order.Weight,
		Price:         order.Price,
		Hash:          hash.GenerateHash(),
		StatusUpdated: time.Now(),
	}
}

func (t *Order) UpdateStatus(newStatus OrderStatus) error {
	switch newStatus {
	case ReturnToDeliverer:
		switch t.Status {
		case ReturnToDeliverer:
			return errors.New("Order was returned to Deliverer early")
		case ReturnFromAddressee:
			return errors.New("The order is a refund now")
		case GivenToAddressee:
			return errors.New("Order was given to Addressee early")
		case InStock:
			if t.ShelfLife.Before(time.Now()) {
				t.Status = ReturnToDeliverer
				t.StatusUpdated = time.Now()
				return nil
			} else {
				return errors.New("The order expire has not ended")
			}
		}
	case ReturnFromAddressee:
		switch t.Status {
		case ReturnToDeliverer:
			return errors.New("Order was returned to Deliverer early")
		case ReturnFromAddressee:
			return errors.New("The order is a refund now")
		case InStock:
			return errors.New("Order on stock now")
		case GivenToAddressee:
			if ((time.Now().Unix() - t.StatusUpdated.Unix()) / 86400) >= ReturnFromAddresseeExpiration {
				return errors.New("Order was given to Addressee more than 2 days")
			}

			t.Status = ReturnFromAddressee
			t.StatusUpdated = time.Now()
			return nil
		}
	case GivenToAddressee:
		switch t.Status {
		case ReturnToDeliverer:
			return errors.New("Order was returned to Deliverer early")
		case ReturnFromAddressee:
			return errors.New("The order is a refund now")
		case GivenToAddressee:
			return errors.New("Order was given to Addressee early")
		case InStock:
			if t.ShelfLife.After(time.Now()) {
				t.Status = GivenToAddressee
				t.StatusUpdated = time.Now()
				return nil
			} else {
				return errors.New("The order expire has ended")
			}
		}
	}

	return errors.New("Undefined status")
}
