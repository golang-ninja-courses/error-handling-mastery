package pipeline

import (
	"context"
	"errors"
	"io"
	"strconv"
)

const maxBatchSize = 4

var (
	errEmptyInput       = errors.New("no data in input")
	errNilChannel       = errors.New("nil channel in input")
	errEmptyTx          = errors.New("input must not contain empty transactions")
	errInvalidBatchSize = errors.New("invalid batch size")
)

type Transaction struct {
	ID int64
}

func (t Transaction) Hash() Hash {
	return newHash([]byte(strconv.FormatInt(t.ID, 10)))
}

type Block struct {
	Hash Hash
}

// Pipeline представляет собой последовательное (но конкуретное внутри) выполнение следующих функций:
//
//  batch -> hashTransactions -> sink -> mergeErrors
//
// Pipeline гарантирует своё завершение только после завершения всех входящих в него функций.
// При этом пайплайн завершается:
//  - при отмене входящего контекста (как следствие того, что от контекста завершатся функции пайплайна);
//  - при получении хотя бы одной ошибки из финального канала ошибок (в таком случае пайплайн требует от функций
//    досрочного завершения и ожидает их);
//  - при безошибочной обработке всех транзакций.
func Pipeline(ctx context.Context, batchSize int, out io.Writer, txs ...Transaction) error {
	// Реализуй меня.
	return nil
}

// batch преобразует входящий слайс транзакций в группы размером batchSize и отправляет в выходной канал:
//  - возвращает ошибку errInvalidBatchSize, если batchSize меньше 1 или больше maxBatchSize;
//  - возвращает ошибку errEmptyInput, если слайс транзакций пуст;
//  - если у очередной транзакции нулевой ID, то функция пишет в выходной канал ошибку errEmptyTx
//    и завершает своё выполнение.
func batch(ctx context.Context, batchSize int, txs ...Transaction) (<-chan []Transaction, <-chan error, error) {
	// Реализуй меня.
	return nil, nil, nil
}

// hashTransactions берёт группы транзакций из входного канала, и считает хеш от группы с помощью CalculateHash:
//  - возвращает ошибку errNilChannel, если на вход получила nil-канал;
//  - если при просчёте хеша возникает ошибка, то функция пишет её в выходной канал и завершает своё выполнение.
func hashTransactions(ctx context.Context, batchc <-chan []Transaction) (<-chan Block, <-chan error, error) {
	// Реализуй меня.
	return nil, nil, nil
}

// sink выводит в out хеши блоков из входного канала в строковой форме, разделяя их через '\n':
//  - возвращает ошибку errNilChannel, если на вход получила nil-канал;
//  - если при записи в out возникает ошибка, то функция пишет её в выходной канал и завершает своё выполнение.
func sink(ctx context.Context, blockc <-chan Block, out io.Writer) (<-chan error, error) {
	// Реализуй меня.
	return nil, nil
}

// mergeErrors сливает все ошибки из входящих каналов в выходной канал:
//  - возвращает ошибку errEmptyInput, если слайс каналов пуст;
//  - возвращает ошибку errNilChannel, если хотя бы один из каналов в слайсе нулевой;
//  - выходной канал закрывается только после того, как будут вычитаны все входные каналы.
func mergeErrors(errcs ...<-chan error) (<-chan error, error) {
	// Реализуй меня.
	return nil, nil
}
