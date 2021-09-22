// Нажмите "Отправить" и задача будет принята.
// Авторское решение можно будет посмотреть в решениях.

package errgroup_fill

import (
	"context"

	_ "golang.org/x/sync/errgroup"
)

type Order struct {
	BuyerName    string
	BuyerAddress string
	SellerName   string
}

func GetOrder(
	ctx context.Context,
	buyerName, buyerAddress, sellerName func(ctx context.Context) (string, error),
) (Order, error) {
	return Order{}, nil
}
