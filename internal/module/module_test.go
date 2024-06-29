package module

import (
	"errors"
	"homework-3/internal/models"
	mock_storage "homework-3/internal/module/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModule(t *testing.T) {
	t.Parallel()

	t.Run("AddOrder", func(t *testing.T) {
		t.Parallel()

		t.Run("ExistanceOrderId", func(t *testing.T) {
			t.Parallel()
			// arrange
			ctrl := gomock.NewController(t)
			mockStorage := mock_storage.NewMockStorage(ctrl)
			module := NewModule(Deps{Storage: mockStorage})
			mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(true, nil)
			// act
			err := module.AddOrder(models.Order{OrderId: 1})
			// assert

			require.NotNil(t, err)
			assert.Equal(t, err.Error(), "Order with this orderId exist in stock")
		})
		t.Run("UnExistanceOrderId", func(t *testing.T) {
			t.Parallel()
			// arrange
			t.Run("ErrorFromStorage", func(t *testing.T) {
				t.Parallel()
				// arrange
				ctrl := gomock.NewController(t)
				mockStorage := mock_storage.NewMockStorage(ctrl)
				module := NewModule(Deps{Storage: mockStorage})
				mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(false, errors.New("Some Error"))
				// act
				err := module.AddOrder(models.Order{OrderId: 1})
				// assert

				require.NotNil(t, err)
				assert.Equal(t, err.Error(), "Some Error")
			})
			t.Run("BadWrap", func(t *testing.T) {
				t.Parallel()
				// arrange
				ctrl := gomock.NewController(t)
				mockStorage := mock_storage.NewMockStorage(ctrl)
				module := NewModule(Deps{Storage: mockStorage})
				mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(false, nil)
				wraper, _ := models.GetWrapper(1)

				// act
				err := module.AddOrder(models.Order{OrderId: 1, Weight: 50, Wrapper: wraper})
				// assert

				require.NotNil(t, err)
				assert.Equal(t, err.Error(), "Ошибка! Вес заказа при упаковке в пакет должен быть меньше 10кг. Вес посылки: 50")
			})
			t.Run("Successful", func(t *testing.T) {
				t.Parallel()
				// arrange
				ctrl := gomock.NewController(t)
				mockStorage := mock_storage.NewMockStorage(ctrl)
				module := NewModule(Deps{Storage: mockStorage})
				mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(false, nil)
				mockStorage.EXPECT().AddOrder(gomock.Any()).Return(nil)
				wraper, _ := models.GetWrapper(1)

				// act
				err := module.AddOrder(models.Order{OrderId: 1, Weight: 9, Wrapper: wraper})
				// assert

				assert.Nil(t, err)
			})
		})
	})
	t.Run("DispatchOrders", func(t *testing.T) {
		t.Parallel()
		t.Run("UnExistanceOrderId", func(t *testing.T) {
			t.Parallel()
			// arrange
			ctrl := gomock.NewController(t)
			mockStorage := mock_storage.NewMockStorage(ctrl)
			module := NewModule(Deps{Storage: mockStorage})
			mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(false, nil)
			ids := []models.OrderId{1, 2}
			// act
			err := module.DispatchOrders(ids)
			// assert
			require.NotNil(t, err)
			assert.Equal(t, err.Error(), "Order with orderId: 1 not exist")
		})
		t.Run("OrderIdWithDifferentAddressee", func(t *testing.T) {
			t.Parallel()
			// arrange
			ctrl := gomock.NewController(t)
			mockStorage := mock_storage.NewMockStorage(ctrl)
			module := NewModule(Deps{Storage: mockStorage})
			mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(true, nil)
			mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(2)).Return(true, nil)
			idsMap := map[models.OrderId]interface{}{1: struct{}{}, 2: struct{}{}}
			ids := []models.OrderId{1, 2}
			mp := map[models.AddresseeId][]models.OrderId{111: {2}, 1232: {1}}
			mockStorage.EXPECT().GetAddresseeIds(idsMap).Return(mp, nil)

			// act
			err := module.DispatchOrders(ids)
			// assert
			require.NotNil(t, err)
		})
		t.Run("Successful", func(t *testing.T) {
			t.Parallel()
			// arrange
			ctrl := gomock.NewController(t)
			mockStorage := mock_storage.NewMockStorage(ctrl)
			module := NewModule(Deps{Storage: mockStorage})
			mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(1)).Return(true, nil)
			mockStorage.EXPECT().CheckExistanceOrder(models.OrderId(2)).Return(true, nil)
			ids := []models.OrderId{2, 1}
			idsMap := map[models.OrderId]interface{}{1: struct{}{}, 2: struct{}{}}
			mp := map[models.AddresseeId][]models.OrderId{111: {1, 2}}
			mockStorage.EXPECT().GetAddresseeIds(idsMap).Return(mp, nil)
			mockStorage.EXPECT().UpdateOrderStatus(models.OrderId(1), gomock.Any()).Return(nil)
			mockStorage.EXPECT().UpdateOrderStatus(models.OrderId(2), gomock.Any()).Return(nil)

			// act
			err := module.DispatchOrders(ids)
			// assert
			require.Nil(t, err)
		})

	})
}
