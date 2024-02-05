package service

import (
	"context"

	"github.com/adhistria/backend/go-vue-project/internal/order/datastore"
	"github.com/adhistria/backend/go-vue-project/internal/order/entity"
)

type OrderService struct {
	orderRepo datastore.OrderRepository
}

func NewOrderService(repo datastore.OrderRepository) OrderService {
	return OrderService{orderRepo: repo}
}

func (o OrderService) Search(ctx context.Context, option entity.Option) (*entity.OrderResponse, error) {
	data, err := o.orderRepo.Search(ctx, option)
	if err != nil {
		return nil, err
	}

	for _, order := range data {
		order.ConvertDateToString()
	}

	orderSummary, err := o.orderRepo.GetTotal(option)
	if err != nil {
		return nil, err
	}

	response := entity.OrderResponse{
		Data:        data,
		TotalRows:   orderSummary.TotalData,
		TotalAmount: orderSummary.TotalAmount,
	}

	return &response, nil
}
