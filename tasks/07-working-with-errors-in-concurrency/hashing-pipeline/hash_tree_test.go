package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
	cases := []struct {
		name        string
		input       []Hashable
		strExpected string
		errExpected error
	}{
		{
			name:        "nil input",
			input:       nil,
			strExpected: "",
			errExpected: errNothingToHash,
		},
		{
			name:        "single input",
			input:       []Hashable{Transaction{ID: 1}},
			strExpected: "488ad59efeafae6af0be932803b65470b908c33377f4ff99aad1fd8c68c4f463",
			errExpected: nil,
		},
		{
			name:        "same two elements in input",
			input:       []Hashable{Transaction{ID: 1}, Transaction{ID: 1}},
			strExpected: "488ad59efeafae6af0be932803b65470b908c33377f4ff99aad1fd8c68c4f463",
			errExpected: nil,
		},
		{
			name:        "same three elements in input",
			input:       []Hashable{Transaction{ID: 1}, Transaction{ID: 1}, Transaction{ID: 1}},
			strExpected: "06852de4b6b79a48f7c7bf7f71fedc1e73d5e79723aeb43bf813858ba72174b8",
			errExpected: nil,
		},
		{
			name:        "different three elements in input",
			input:       []Hashable{Transaction{ID: 1}, Transaction{ID: 2}, Transaction{ID: 3}},
			strExpected: "f981662b1dcd91b2569a56fce8c590b04bc062ee22d459e49bc507638c8099a2",
			errExpected: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			h, err := CalculateHash(tt.input)

			assert.ErrorIs(t, tt.errExpected, err)
			assert.Equal(t, tt.strExpected, h.String())
		})
	}
}
