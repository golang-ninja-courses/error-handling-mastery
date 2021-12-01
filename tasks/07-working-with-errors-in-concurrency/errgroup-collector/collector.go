package collector

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

const maxSectors = 10

var ErrTooMuchSectors = errors.New("too much sectors")

type SensorValue struct {
	SensorID string
	Value    float64
}

type Sector interface {
	ID() string
	GetSensorValues(ctx context.Context) ([]SensorValue, error)
}

// Collect конкурентно собирает значения датчиков с секторов, объединяет их в один
// слайс данных и выдаёт наружу.
// При превышении лимита на количество секторов возвращает ошибку ErrTooMuchSectors.
// При возникновении ошибки во время опроса очередного сектора, функция завершает
// свою работу и возвращает эту ошибку.
func Collect(ctx context.Context, sectors []Sector) ([]SensorValue, error) {
	// Для сохранения импортов. Удали эту строку.
	_ = errgroup.Group{}

	// Реализуй меня.
	return nil, nil
}
