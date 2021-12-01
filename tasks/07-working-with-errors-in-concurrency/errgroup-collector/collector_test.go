package collector

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

const defaultCollectTimeout = 5 * time.Second

func TestCollect(t *testing.T) {
	errSensorUnavailable := errors.New("sensor is unavailable")

	cases := []struct {
		name             string
		sectors          []Sector
		collectTimeout   time.Duration
		expectedErr      error
		expectedExecTime time.Duration // Всегда берём с запасом из-за погрешности на среду выполнения.
		expectedValues   []SensorValue
	}{
		{
			name:           "no sectors no problems 1",
			sectors:        nil,
			expectedValues: []SensorValue{},
		},
		{
			name:           "no sectors no problems 2",
			sectors:        []Sector{},
			expectedValues: []SensorValue{},
		},
		{
			name: "too much sectors",
			sectors: []Sector{
				newSimpleSector("1"), newSimpleSector("2"), newSimpleSector("3"),
				newSimpleSector("4"), newSimpleSector("5"), newSimpleSector("6"),
				newSimpleSector("7"), newSimpleSector("8"), newSimpleSector("9"),
				newSimpleSector("1"), newSimpleSector("10"), newSimpleSector("11"),
			},
			expectedErr: ErrTooMuchSectors,
		},
		{
			name: "one fast sector",
			sectors: []Sector{
				sectorMock{values: newSensorValues(1, 3)},
			},
			expectedValues: newSensorValues(1, 3),
		},
		{
			name: "one slow sector",
			sectors: []Sector{
				sectorMock{values: newSensorValues(1, 3), pollingTime: time.Second},
			},
			expectedExecTime: time.Second * 2,
			expectedValues:   newSensorValues(1, 3),
		},
		{
			name: "three sectors with the same data",
			sectors: []Sector{
				sectorMock{values: newSensorValues(1, 3), pollingTime: time.Second / 2},
				sectorMock{values: newSensorValues(1, 3), pollingTime: time.Second / 2},
				sectorMock{values: newSensorValues(1, 3), pollingTime: time.Second / 2},
			},
			expectedExecTime: time.Second * 2,
			expectedValues: []SensorValue{
				sv("1", 1), sv("2", 2), sv("3", 3),
				sv("1", 1), sv("2", 2), sv("3", 3),
				sv("1", 1), sv("2", 2), sv("3", 3),
			},
		},
		{
			name: "ten sectors with diff polling time",
			sectors: []Sector{
				sectorMock{id: "1", values: newSensorValues(1, 3), pollingTime: time.Second},
				sectorMock{id: "2", values: newSensorValues(4, 6), pollingTime: time.Second / 2},
				sectorMock{id: "3", values: newSensorValues(7, 9), pollingTime: time.Second / 3},
				sectorMock{id: "4", values: newSensorValues(10, 12), pollingTime: time.Second},
				sectorMock{id: "5", values: newSensorValues(13, 15), pollingTime: time.Second / 2},
				sectorMock{id: "6", values: newSensorValues(16, 18), pollingTime: time.Second / 3},
				sectorMock{id: "7", values: newSensorValues(19, 21), pollingTime: time.Second},
				sectorMock{id: "8", values: newSensorValues(22, 24), pollingTime: time.Second / 2},
				sectorMock{id: "9", values: newSensorValues(25, 27), pollingTime: time.Second / 3},
				sectorMock{id: "10", values: newSensorValues(28, 30), pollingTime: time.Second},
			},
			expectedExecTime: 2 * time.Second,
			expectedValues:   newSensorValues(1, 30),
		},
		{
			name: "poll sensors error",
			sectors: []Sector{
				sectorMock{id: "1", values: newSensorValues(1, 3), pollingTime: defaultCollectTimeout},
				sectorMock{id: "2", values: newSensorValues(4, 6), pollingTime: defaultCollectTimeout},
				sectorMock{id: "3", err: errSensorUnavailable, pollingTime: time.Second},
				sectorMock{id: "4", values: newSensorValues(10, 12), pollingTime: defaultCollectTimeout},
			},
			expectedExecTime: 2 * time.Second,
			expectedErr:      errSensorUnavailable,
		},
		{
			name: "collect timeout",
			sectors: []Sector{
				sectorMock{id: "1", values: newSensorValues(1, 3), pollingTime: time.Second},
				sectorMock{id: "2", values: newSensorValues(4, 6), pollingTime: time.Second / 2},
				sectorMock{id: "3", values: newSensorValues(7, 9), pollingTime: defaultCollectTimeout * 2},
				sectorMock{id: "4", values: newSensorValues(10, 12), pollingTime: defaultCollectTimeout},
			},
			collectTimeout:   time.Second,
			expectedExecTime: 2 * time.Second,
			expectedErr:      context.DeadlineExceeded,
		},
		{
			name: "large amount of data",
			sectors: []Sector{
				sectorMock{values: newSensorValues(1, 2000)},
				sectorMock{values: newSensorValues(2001, 4000)},
				sectorMock{values: newSensorValues(4001, 6000)},
			},
			expectedValues: newSensorValues(1, 6000),
		},
		{
			name: "error while large amount of data",
			sectors: []Sector{
				sectorMock{values: newSensorValues(1, 2000)},
				sectorMock{err: errSensorUnavailable, pollingTime: time.Second},
				sectorMock{values: newSensorValues(4001, 6000)},
			},
			expectedErr: errSensorUnavailable,
		},
	}

	type result struct {
		err    error
		values []SensorValue
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			defer goleak.VerifyNone(t)

			timeout := tt.collectTimeout
			if timeout == 0 {
				timeout = defaultCollectTimeout
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			resultc := make(chan result, 1)

			go func() {
				values, err := Collect(ctx, tt.sectors)
				resultc <- result{err, values}
			}()

			start := time.Now()

			select {
			case <-time.After(defaultCollectTimeout):
				t.Fatal("tested code is too slow: Collect blocked?")

			case r := <-resultc:
				require.ErrorIs(t, r.err, tt.expectedErr)

				if tt.expectedExecTime != 0 {
					assert.LessOrEqual(t, time.Since(start), tt.expectedExecTime,
						"sectors did not collected in parallel")
				}
				assert.ElementsMatch(t, r.values, tt.expectedValues)
			}
		})
	}
}

var _ Sector = sectorMock{}

type sectorMock struct {
	id          string
	err         error
	values      []SensorValue
	pollingTime time.Duration
}

func (s sectorMock) ID() string {
	return s.id
}

func (s sectorMock) GetSensorValues(ctx context.Context) ([]SensorValue, error) {
	if s.pollingTime != 0 {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(s.pollingTime):
		}
	}
	return s.values, s.err
}

func newSimpleSector(id string) Sector {
	return sectorMock{id: id, values: newSensorValues(1, 3), pollingTime: time.Second}
}

func newSensorValues(from, to int) []SensorValue {
	if from > to {
		panic(fmt.Sprintf("invalid newSensorValues call: %d > %d", from, to))
	}

	vals := make([]SensorValue, 0, to-from)
	for i := from; i <= to; i++ {
		vals = append(vals, sv(strconv.Itoa(i), float64(i)))
	}
	return vals
}

func sv(id string, v float64) SensorValue {
	return SensorValue{
		SensorID: id,
		Value:    v,
	}
}
