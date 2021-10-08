package orders

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Order struct {
	BuyerName    string
	BuyerAddress string
	SellerName   string
}

type StringValueGetter func(ctx context.Context) (string, error)

func GetOrder(ctx context.Context, getBuyerName, getBuyerAddress, getSellerName StringValueGetter) (Order, error) {
	// Для сохранения импортов. Удали эту строку.
	_ = errgroup.Group{}

	// Реализуй меня.
	return Order{}, nil
}
