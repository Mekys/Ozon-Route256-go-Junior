//go:generate mockgen -source ./module.go -destination=./mocks/storage.go -package=mock_storage

package module

import (
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/models"
	"homework-3/internal/storage"
)

type Storage interface {
	AddOrder(models.Order) error
	CheckExistanceOrder(models.OrderId) (bool, error)
	UpdateOrderStatus(models.OrderId, storage.OrderStatus) error
	GetAddresseeIds(orderIds map[models.OrderId]interface{}) (map[models.AddresseeId][]models.OrderId, error)
	ListOrders(models.AddresseeId) ([]models.Order, error)
	ListRefund(int64, int64) ([]models.Order, error)
}

type Deps struct {
	Storage Storage
}

type Module struct {
	Deps
}

// NewModule .. TODO сделать описание функции
func NewModule(d Deps) Module {
	return Module{Deps: d}
}

func (m Module) AddOrder(order models.Order) error {

	existance, err := m.Storage.CheckExistanceOrder(order.OrderId)

	if err != nil {
		return err
	}

	if order.Wrapper != nil { //Если у нас есть какая-то упаковка, то нужно добавить за нее стоимость
		err = order.Wrapper.AddPriceForWrap(&order)

		if err != nil {
			return err
		}
	}

	if existance {
		return errors.New("Order with this orderId exist in stock")
	}

	// Запись в хранилище
	return m.Storage.AddOrder(order)
}

func (m Module) ReturnToDeliverer(orderId models.OrderId) error {
	if err := m.Storage.UpdateOrderStatus(orderId, storage.ReturnToDeliverer); err != nil {
		return err
	} else {
		fmt.Println("Successful return to deliverer")
		return nil
	}
}
func (m Module) ReturnOrder(order models.Order) error {
	if err := m.Storage.UpdateOrderStatus(order.OrderId, storage.ReturnFromAddressee); err != nil {
		return err
	} else {
		fmt.Println("Successful return from Addressee")
		return nil
	}
}
func (m Module) ListRefund(pageLen int64, pageNumber int64) error {
	refund, err := m.Storage.ListRefund(pageLen, pageNumber)
	if err != nil {
		return err
	}
	fmt.Printf("Refund from %d to %d\n", (pageNumber-1)*pageLen+1, pageNumber*pageLen)
	fmt.Println("_____________________________________________")
	if len(refund) == 0 {
		fmt.Println("None")
	} else {
		for index, item := range refund {
			fmt.Printf("[%d] OrderId: %d AddresseeId: %d ShelfLife: %s \n", (int(pageNumber)-1)*int(pageLen)+1+index, item.OrderId, item.AddresseeId, item.ShelfLife.Format("2006-01-02"))
		}
	}

	fmt.Println("______________end of list______________________")
	return nil
}
func (m Module) ListOrders(addresseeId int64, count int64) error {
	orders, err := m.Storage.ListOrders(models.AddresseeId(addresseeId))
	if err != nil {
		return err
	}
	for i, j := 0, len(orders)-1; i < j; i, j = i+1, j-1 {
		orders[i], orders[j] = orders[j], orders[i]
	}
	if count == 0 {
		fmt.Printf("All orders with AddresseeId: %d \n", addresseeId)
	} else {
		fmt.Printf("Last %d orders with AddresseeId: %d \n", count, addresseeId)
		orders = orders[:min(int(count), cap(orders))]
	}
	fmt.Println("_____________________________________________")
	if len(orders) == 0 {
		fmt.Println("None")
	} else {
		for index, item := range orders {
			fmt.Printf("[%d] OrderId: %d AddresseeId: %d ShelfLife: %s \n", index+1, item.OrderId, item.AddresseeId, item.ShelfLife.Format("2006-01-02"))
		}
	}

	fmt.Println("______________end of list______________________")
	return nil
}

func (m Module) DispatchOrders(orderIds []models.OrderId) error {
	orderIdMap := make(map[models.OrderId]interface{})

	for _, value := range orderIds {
		isExist, err := m.Storage.CheckExistanceOrder(value)

		if err != nil {
			return err
		}

		if !isExist {
			return fmt.Errorf("Order with orderId: %d not exist", value)
		}

		orderIdMap[value] = struct{}{}
	}

	addresseeIds, err := m.Storage.GetAddresseeIds(orderIdMap)

	if err != nil {
		return err
	}

	if len(addresseeIds) > 1 {
		bs, _ := json.Marshal(addresseeIds)
		return fmt.Errorf("Order have not same Adressee\nDetails:{AddresseeId_1:[OrderId_1, OrderId_2], AddresseeId_2:[OrderId_3, OrderId_4]}\n %s", bs)
	}

	for key, _ := range orderIdMap {
		if err := m.Storage.UpdateOrderStatus(key, storage.GivenToAddressee); err != nil {
			return fmt.Errorf("OrderId: %d %w", key, err)
		}
	}

	fmt.Println("Successful dispatch orders")
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
