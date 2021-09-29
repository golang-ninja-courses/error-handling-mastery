package pipeline

import (
	"context"
	"errors"
	_ "fmt"
	"strconv"
	_ "sync"
)

type Transaction struct {
	ID int64
}

func (t Transaction) Hash() Hash {
	return hash([]byte(strconv.FormatInt(t.ID, 10))[:])
}

type Block struct {
	Hash Hash
}

const (
	maxBatchSize = 4
)

var (
	errEmptyInput       = errors.New("input strings couldn't be empty")
	errEmptyTx          = errors.New("input shouldn't contain empty transactions")
	errInvalidBatchSize = errors.New("invalid batch size")
)

// Source преобразует входящий слайс транзакций input в outbound-канал.
// - возвращает ошибку errEmptyInput, если input пуст;
// - пишет в outbound-канал ошибку errEmptyTx, если у транзакции input нулевой ID и завершает своё выполнение.
func Source(ctx context.Context, input ...Transaction) (<-chan Transaction, <-chan error, error) {
	// Реализовать.
	return nil, nil, nil
}

// Aggregate группирует транзакции из inbound-канала in в батчи размером batchSize и отправляет в outbound-канал.
// - возвращает ошибку errInvalidBatchSize, если batchSize меньше 1 или больше maxBatchSize.
func Aggregate(ctx context.Context, batchSize int, in <-chan Transaction) (<-chan []Transaction, <-chan error, error) {
	// Реализовать.
	return nil, nil, nil
}

// HashTxs берет батчи транзакций из inbound-канала, хеширует функцией sha265 и отправляет получившийся блок
// с общим хешем в outbound-канал.
// Принцип хеширования визуализирован в поясняющей диаграмме в шаге. Реализуйте его в функции CalculateHash.
// Ошибку от CalculateHash нужно писать в outbound-канал и завершать выполнение.
func HashTxs(ctx context.Context, in <-chan []Transaction) (<-chan Block, <-chan error, error) {
	// Реализовать.
	return nil, nil, nil
}

// Sink выводит на экран хеши блоков из inbound-канала.
func Sink(_ context.Context, in <-chan Block) (<-chan error, error) {
	// Реализовать.
	return nil, nil
}

// Merge принимает на вход слайс каналов и возвращает один канал.
// Задача функции – всё, что есть во входящих inbound-каналах, слить в один outbound-канал.
func Merge(errcs ...<-chan error) <-chan error {
	// Реализовать.
	return nil
}

// Pipeline вишенка на торте, функция, которая прогонит пайплайн.
// - пайплайн таков: Source -> Aggregate -> HashTxs -> Sink.
// - в конце Merge объединяет ошибки из каналов каждого этапа, и если там что-то есть то возвращает эту ошибку.
// - в Pipeline надо рулить контекстом, чтобы при возникновении ошибки на одном из этапов,
// остальные досрочно завершали свою работу.
func Pipeline(ctx context.Context, txs ...Transaction) error {
	// Реализовать.
	return nil
}
