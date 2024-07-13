package service

import (
	"context"
	"homework-3/internal/models"
	"homework-3/internal/module"
	__order "homework-3/pkg/api/proto/order/v1/order/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TelephoneService struct {
	Module module.Module
	__order.UnimplementedOrderServer
}

func (ts *TelephoneService) AddOrder(_ context.Context, order *__order.AddOrderRequest) (*emptypb.Empty, error) {
	wrapper, _ := models.GetWrapper(int(*order.WrapType))
	return nil, ts.Module.AddOrder(
		models.Order{
			OrderId:     models.OrderId(order.OrderId),
			AddresseeId: models.AddresseeId(order.AddresseeId),
			Weight:      int(*order.Weight),
			Price:       int(*order.Price),
			Wrapper:     wrapper,
			ShelfLife:   order.ShelfLife.AsTime(),
		})
}
func (ts *TelephoneService) ReturnToDeliverer(_ context.Context, order *__order.ReturnToDelivererRequest) (*emptypb.Empty, error) {
	return nil, ts.Module.ReturnToDeliverer(models.OrderId(order.OrderId))
}
func (*TelephoneService) ListOrder(context.Context, *__order.ListOrderRequest) (*__order.ListOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (ts *TelephoneService) GiveToAddressee(_ context.Context, ord *__order.GiveToAddresseeRequest) (*emptypb.Empty, error) {
	return nil, ts.Module.ReturnOrder(models.Order{OrderId: models.OrderId(ord.OrderId), AddresseeId: models.AddresseeId(ord.AddresseeId)})
}
func (ts *TelephoneService) ReturnFromAddressee(_ context.Context, order *__order.ReturnFromAddresseeRequest) (*emptypb.Empty, error) {
	return nil, ts.Module.DispatchOrders(GetArrayOfOrderId(order.OrderIds))
}
func (*TelephoneService) ListRefund(context.Context, *__order.ListRefundRequest) (*__order.ListRefundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRefund not implemented")
}

func GetArrayOfOrderId(OrderIds []int64) []models.OrderId {
	result := make([]models.OrderId, 0)

	for value := range OrderIds {
		result = append(result, models.OrderId(value))
	}

	return result
}
