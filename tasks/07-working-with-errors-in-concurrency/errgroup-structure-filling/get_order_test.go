package orders

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testTimeout = 100 * time.Millisecond

	buyerName    = "John Doe"
	sellerName   = "Jane Doe"
	buyerAddress = "123 Maple Street. Anytown, PA 17101"
)

var (
	buyerNameFunc    = func(ctx context.Context) (string, error) { return buyerName, nil }
	buyerAddressFunc = func(ctx context.Context) (string, error) { return buyerAddress, nil }
	sellerNameFunc   = func(ctx context.Context) (string, error) { return sellerName, nil }

	slowBuyerNameFunc = func(ctx context.Context) (string, error) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(10 * testTimeout):
			return buyerName, nil
		}
	}

	errNotFound = errors.New("not found error")

	errFunc = func(ctx context.Context) (string, error) {
		return "", errNotFound
	}

	filledOrder = Order{
		BuyerName:    buyerName,
		BuyerAddress: buyerAddress,
		SellerName:   sellerName,
	}
)

func TestGetOrder(t *testing.T) {
	cases := []struct {
		name             string
		buyerNameFunc    StringValueGetter
		buyerAddressFunc StringValueGetter
		sellerNameFunc   StringValueGetter
		orderExpected    Order
		errExpected      error
		ctxErrExpected   error
	}{
		{
			name:             "no errors",
			buyerNameFunc:    buyerNameFunc,
			buyerAddressFunc: buyerAddressFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    filledOrder,
		},
		{
			name:             "buyer name error",
			buyerNameFunc:    errFunc,
			buyerAddressFunc: buyerAddressFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    Order{},
			errExpected:      errNotFound,
			ctxErrExpected:   nil,
		},
		{
			name:             "slow buyer name",
			buyerNameFunc:    slowBuyerNameFunc,
			buyerAddressFunc: buyerAddressFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    Order{},
			errExpected:      context.DeadlineExceeded,
			ctxErrExpected:   context.DeadlineExceeded,
		},
		{
			name:             "slow buyer name and error",
			buyerNameFunc:    slowBuyerNameFunc,
			buyerAddressFunc: errFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    Order{},
			errExpected:      errNotFound,
			ctxErrExpected:   nil,
		},
		{
			name: "total functions execution time greater than test timeout",
			buyerNameFunc: func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(testTimeout / 2):
					return buyerName, nil
				}
			},
			buyerAddressFunc: func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(testTimeout / 2):
					return buyerAddress, nil
				}
			},
			sellerNameFunc: func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(testTimeout / 2):
					return sellerName, nil
				}
			},
			orderExpected: filledOrder,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
			defer cancel()

			order, err := GetOrder(ctx, tt.buyerNameFunc, tt.buyerAddressFunc, tt.sellerNameFunc)
			require.ErrorIs(t, err, tt.errExpected)
			assert.Equal(t, tt.orderExpected, order)

			assert.ErrorIs(t, ctx.Err(), tt.ctxErrExpected)
		})
	}
}
