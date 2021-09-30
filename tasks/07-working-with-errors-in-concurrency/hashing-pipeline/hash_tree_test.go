package pipeline

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
	type tCase struct {
		name         string
		input        []Hashable
		hashExpected string
		errExpected  error
	}

	cases := []tCase{
		{
			name:        "nil input",
			input:       nil,
			errExpected: errNothingToHash,
		},
		{
			name:        "empty input",
			input:       []Hashable{},
			errExpected: errNothingToHash,
		},
		{
			name:         "single input",
			input:        []Hashable{Transaction{ID: 1}},
			hashExpected: "d829bfd6219e2dab703ff954799e50e8ab9d8377490b69691eb0c0c5fcdb2ec4",
		},
		{
			name:         "same two elements", // Обратите внимание на равенство хешу из теста выше.
			input:        []Hashable{Transaction{ID: 1}, Transaction{ID: 1}},
			hashExpected: "d829bfd6219e2dab703ff954799e50e8ab9d8377490b69691eb0c0c5fcdb2ec4",
		},
		{
			name:         "different two elements",
			input:        []Hashable{Transaction{ID: 1}, Transaction{ID: 2}},
			hashExpected: "7de236613dd3d9fa1d86054a84952f1e0df2f130546b394a4d4dd7b76997f607",
		},
		{
			name:         "different three elements",
			input:        []Hashable{Transaction{ID: 1}, Transaction{ID: 2}, Transaction{ID: 3}},
			hashExpected: "5b7534123197114fa7e7459075f39d89ffab74b5c3f31fad48a025b931ff5a01",
			errExpected:  nil,
		},
		{
			name:         "same three elements",
			input:        []Hashable{Transaction{ID: 1}, Transaction{ID: 1}, Transaction{ID: 1}},
			hashExpected: "599420fae07766d092aeee690e1940183001db1ebbb50b3a606e16d2168a0f05",
			errExpected:  nil,
		},
		{
			name:         "same four elements", // Обратите внимание на равенство хешу из теста выше.
			input:        []Hashable{Transaction{ID: 1}, Transaction{ID: 1}, Transaction{ID: 1}, Transaction{ID: 1}},
			hashExpected: "599420fae07766d092aeee690e1940183001db1ebbb50b3a606e16d2168a0f05",
			errExpected:  nil,
		},
		{
			name:         "different four elements",
			input:        []Hashable{Transaction{ID: 1}, Transaction{ID: 2}, Transaction{ID: 3}, Transaction{ID: 4}},
			hashExpected: "8b186d4723474e69fd14c28384063e2031d5da66b97844d5973a9e9bf7dcfeeb",
			errExpected:  nil,
		},
	}

	const n = 10_000_000
	cases = append(cases, tCase{
		name:         fmt.Sprintf("different %d elements", n),
		input:        makeHashableFromTransactions(n),
		hashExpected: "d5844190df77ecd8d6c512f040ecab1b725bb28867a8d87b44af1ce7735b6a7d",
	})

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			h, err := CalculateHash(tt.input)

			require.ErrorIs(t, tt.errExpected, err)
			assert.Equal(t, tt.hashExpected, h.String())
		})
	}
}

func TestCalculateHash_DoNotAffectInput(t *testing.T) {
	in := makeHashableFromTransactions(100)
	for i := 0; i < 10; i++ {
		h, err := CalculateHash(in)
		require.NoError(t, err)
		assert.Equal(t, "3f9ff7574a96d08662a7b9826a320f490e7a13e6d8a5639f2203d613cf082298", h.String())
	}
}

func makeHashableFromTransactions(n int) []Hashable {
	txs := make([]Hashable, n)
	for i := 0; i < n; i++ {
		txs[i] = Transaction{ID: int64(i)}
	}
	return txs
}
