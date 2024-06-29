package storage

import (
	"context"
	"errors"
	"homework-3/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	p *pgxpool.Pool
}

const (
	orderTable = "orders"
)

var (
	orderColumns = []string{"orderid", "addresseeid", "shelflife", "order_status", "status_updated_date", "hash_code", "weight", "price"}
)

func NewStorage(currPool *pgxpool.Pool) Storage {
	return Storage{p: currPool}
}

func (s Storage) AddOrder(order models.Order) error {

	newValue := transform(order)

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert(orderTable).Columns(orderColumns...).Values(
		newValue.OrderId,
		newValue.AddresseeId,
		newValue.ShelfLife,
		newValue.Status,
		newValue.StatusUpdated,
		newValue.Hash,
		newValue.Weight,
		newValue.Price,
	)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = s.p.Exec(context.Background(),
		rawQuery,
		args...)

	return err
}

func (s Storage) CheckExistanceOrder(orderId models.OrderId) (bool, error) {

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(orderColumns[0]).
		From(orderTable).
		Where("orderId = $1", orderId).
		Limit(1)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	var ids int64 = -1
	err = s.p.QueryRow(context.Background(), rawQuery, args...).Scan(&ids)
	if err != nil {
		return false, nil
	}

	//orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[orderRecord])

	// if err != nil {
	// 	return false, err
	// }

	return (ids != -1), nil

}

// случайно сломал этот метод, когда переписывал и не успел разобраться
func (s Storage) UpdateOrderStatus(orderId models.OrderId, newStatus OrderStatus) error {

	isOrderExist, err := s.CheckExistanceOrder(orderId)

	if err != nil {
		return err
	}

	if !isOrderExist {
		return errors.New("Order with this orderId not in stock")
	}

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(orderColumns...).
		From(orderTable).
		Where("orderid = $1", orderId).
		Limit(1)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	var order []orderRecord
	err = pgxscan.Select(context.Background(), s.p, &order, rawQuery, args...)
	if err != nil {
		return err
	}
	updateStatusErr := order[0].UpdateStatus(newStatus)

	if updateStatusErr != nil {
		return updateStatusErr
	}

	queryUpdate := sq.
		Update(orderTable).
		Set("order_status", newStatus).
		Where(sq.Eq{"orderId": orderId}).
		PlaceholderFormat(sq.Dollar)

	rawQueryUpdate, argsUpdate, err := queryUpdate.ToSql()
	if err != nil {
		return err
	}

	_, err = s.p.Query(context.Background(), rawQueryUpdate, argsUpdate...)

	if err != nil {
		return err
	}

	return nil
}
func (s Storage) ListRefund(pageLen int64, numberPage int64) ([]models.Order, error) {

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("orderId", "addresseeId", "shelfLife").
		From(orderTable).
		Where("order_status = $1", ReturnFromAddressee).
		Offset(uint64(numberPage-1) * uint64(pageLen)).
		Limit(uint64(pageLen))

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var refund []models.Order

	err = pgxscan.Select(context.Background(), s.p, &refund, rawQuery, args...)

	if err != nil {
		return nil, err
	}

	return refund, nil

}
func (s Storage) ListOrders(addresseeId models.AddresseeId) ([]models.Order, error) {

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("orderid", "addresseeid", "shelflife").
		From(orderTable).
		Where(sq.Eq{"order_status": InStock, "addresseeId": addresseeId})

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var orders []models.Order

	err = pgxscan.Select(context.Background(), s.p, &orders, rawQuery, args...)

	if err != nil {
		return nil, err
	}

	return orders, nil
}
func (s Storage) GetAddresseeIds(orderIds map[models.OrderId]interface{}) (map[models.AddresseeId][]models.OrderId, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(orderColumns...).
		From(orderTable)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var records []orderRecord

	err = pgxscan.Select(context.Background(), s.p, &records, rawQuery, args...)
	orderIdsGroupByAddresseeId := make(map[models.AddresseeId][]models.OrderId)
	for _, existanceOrder := range records {
		if _, ok := orderIds[existanceOrder.OrderId]; ok {
			if value, isFindBefore := orderIdsGroupByAddresseeId[existanceOrder.AddresseeId]; isFindBefore {
				orderIdsGroupByAddresseeId[existanceOrder.AddresseeId] = append(value, existanceOrder.OrderId)
			} else {
				orderIdsGroupByAddresseeId[existanceOrder.AddresseeId] = append(value, existanceOrder.OrderId)
			}
		}
	}

	return orderIdsGroupByAddresseeId, nil
}
