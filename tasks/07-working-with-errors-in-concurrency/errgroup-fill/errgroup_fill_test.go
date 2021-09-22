package errgroup_fill

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testTimeout       = 100 * time.Millisecond
	funcExecutionTime = 50 * time.Millisecond

	buyerName    = "John Doe"
	sellerName   = "Jane Doe"
	buyerAddress = "123 Maple Street. Anytown, PA 17101"
)

var (
	buyerNameFunc = func(ctx context.Context) (string, error) {
		return buyerName, nil
	}
	buyerAddressFunc = func(ctx context.Context) (string, error) {
		return buyerAddress, nil
	}
	sellerNameFunc = func(ctx context.Context) (string, error) {
		return sellerName, nil
	}

	slowBuyerNameFunc = func(ctx context.Context) (string, error) {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(10 * testTimeout):
			return buyerName, nil
		}
	}

	errFunc = func(ctx context.Context) (string, error) {
		return "", errAny
	}

	order = Order{
		BuyerName:    buyerName,
		BuyerAddress: buyerAddress,
		SellerName:   sellerName,
	}

	errAny = errors.New("any error")
)

func TestGetOrder(t *testing.T) {
	cases := []struct {
		name             string
		buyerNameFunc    func(context.Context) (string, error)
		buyerAddressFunc func(context.Context) (string, error)
		sellerNameFunc   func(context.Context) (string, error)
		orderExpected    Order
		errExpected      error
		checkCtxErrorNil bool
	}{
		{
			name:             "no errors",
			buyerNameFunc:    buyerNameFunc,
			buyerAddressFunc: buyerAddressFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    order,
		},
		{
			name:             "buyer name error",
			buyerNameFunc:    errFunc,
			buyerAddressFunc: buyerAddressFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    order,
			errExpected:      errAny,
		},
		{
			name:             "slow buyer name",
			buyerNameFunc:    slowBuyerNameFunc,
			buyerAddressFunc: buyerAddressFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    order,
			errExpected:      context.DeadlineExceeded,
		},
		{
			name:             "slow buyer name and error",
			buyerNameFunc:    slowBuyerNameFunc,
			buyerAddressFunc: errFunc,
			sellerNameFunc:   sellerNameFunc,
			orderExpected:    order,
			errExpected:      errAny,
			checkCtxErrorNil: true,
		},
		{
			name: "total functions execution time greater than test timeout",
			buyerNameFunc: func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(funcExecutionTime):
					return buyerName, nil
				}
			},
			buyerAddressFunc: func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(funcExecutionTime):
					return buyerAddress, nil
				}
			},
			sellerNameFunc: func(ctx context.Context) (string, error) {
				select {
				case <-ctx.Done():
					return "", ctx.Err()
				case <-time.After(funcExecutionTime):
					return sellerName, nil
				}
			},
			orderExpected: order,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
			defer cancel()

			order, err := GetOrder(ctx, tt.buyerNameFunc, tt.buyerAddressFunc, tt.sellerNameFunc)

			if tt.errExpected == nil {
				assert.Equal(t, tt.orderExpected, order)
			}
			if tt.checkCtxErrorNil {
				assert.Nil(t, ctx.Err())
			}
			assert.ErrorIs(t, tt.errExpected, err)
		})
	}
}
