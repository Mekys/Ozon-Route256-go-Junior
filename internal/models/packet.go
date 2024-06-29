package models

import (
	"errors"
	"fmt"
)

type Wrapper interface {
	AddPriceForWrap(orderForCalc *Order) error
}

type Package struct{}
type Box struct{}
type Tape struct{}

func (pack Package) AddPriceForWrap(orderForCalc *Order) error {

	if orderForCalc.Weight >= 10 {
		return fmt.Errorf("Ошибка! Вес заказа при упаковке в пакет должен быть меньше 10кг. Вес посылки: %d", orderForCalc.Weight)
	}
	orderForCalc.Price += 5

	return nil
}

func (pack Box) AddPriceForWrap(orderForCalc *Order) error {

	if orderForCalc.Weight >= 30 {
		return fmt.Errorf("Ошибка! Вес заказа при упаковке в коробку должен быть меньше 30кг. Вес посылки: %d", orderForCalc.Weight)
	}
	orderForCalc.Price += 20

	return nil
}

func (pack Tape) AddPriceForWrap(orderForCalc *Order) error {

	orderForCalc.Price += 1
	return nil
}

func GetWrapper(idType int) (Wrapper, error) {
	switch idType {
	case 0:
		return nil, nil
	case 1:
		return Package{}, nil
	case 2:
		return Box{}, nil
	case 3:
		return Tape{}, nil
	default:
		return nil, errors.New("Неизвестный тип упаковки")
	}
}
