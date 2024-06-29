package test

import (
	"homework-3/internal/models"
	"homework-3/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	err := storage.AddOrder(models.Order{OrderId: 2, AddresseeId: 4, Price: 10, Weight: 50})
	// assert
	require.NoError(t, err)
}
func TestCheckExistanceOrder(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	isExist, err := storage.CheckExistanceOrder(models.OrderId(1))
	// assert
	require.NoError(t, err)
	assert.True(t, isExist)
}
func TestCheckUnExistanceOrder(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	isExist, err := storage.CheckExistanceOrder(models.OrderId(11))
	// assert
	require.NoError(t, err)
	assert.False(t, isExist)
}
func TestGetSameAddresseeIds(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	orders := map[models.OrderId]interface{}{1: struct{}{}, 4: struct{}{}}
	except := map[models.AddresseeId][]models.OrderId{
		1: {1, 4},
	}
	defer Db.TearDown(t)
	// act
	groupedOrders, err := storage.GetAddresseeIds(orders)
	// assert
	require.NoError(t, err)
	assert.Equal(t, except, groupedOrders)
}
func TestGetDifferentAddresseeIds(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	orders := map[models.OrderId]interface{}{1: struct{}{}, 2: struct{}{}}
	except := map[models.AddresseeId][]models.OrderId{
		1: {1},
		2: {2},
	}
	defer Db.TearDown(t)
	// act
	groupedOrders, err := storage.GetAddresseeIds(orders)
	// assert
	require.NoError(t, err)
	assert.Equal(t, except, groupedOrders)
}
func TestGetOrders(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	except := []models.OrderId{2}
	actual := make([]models.OrderId, 1)
	defer Db.TearDown(t)
	// act
	orders, err := storage.ListOrders(models.AddresseeId(2))
	// assert
	require.NoError(t, err)
	require.Len(t, orders, 1)
	actual[0] = orders[0].OrderId
	assert.Equal(t, except, actual)
}
func TestGetRefund(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	orders, err := storage.ListRefund(10, 1)
	// assert
	require.NoError(t, err)
	require.Len(t, orders, 6)
}
func TestGetRefundEmpty(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	orders, err := storage.ListRefund(10, 100)
	// assert
	require.NoError(t, err)
	require.Len(t, orders, 0)
}
func TestUpdateOrderGiveToAddresseeFail(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	err := storage.UpdateOrderStatus(models.OrderId(1), 1)
	// assert
	require.Error(t, err)
	assert.Equal(t, err.Error(), "The order expire has ended")
}
func TestUpdateOrderGiveToAddresseeSucces(t *testing.T) {
	// arrange
	Db.SetUp(t, "orders")
	Db.Fill()
	storage := storage.NewStorage(Db.Pool)
	defer Db.TearDown(t)
	// act
	err := storage.UpdateOrderStatus(models.OrderId(2), 1)
	// assert
	require.Nil(t, err)
}
