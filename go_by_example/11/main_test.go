package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntMin(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{-1, 1, -1},
	}

	for _, tt := range tests {
		got := IntMin(tt.a, tt.b)
		assert.Equal(t, tt.want, got)
	}
}

func TestDiv(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := Div(10, 2)
		require.NoError(t, err)
		assert.Equal(t, 5, res)
	})

	t.Run("divide by zero", func(t *testing.T) {
		res, err := Div(10, 0)
		require.Error(t, err)
		assert.True(t, errors.Is(err, ErrDivideByZero))
		assert.Equal(t, 0, res)
	})
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

//go test -v
//go test -v -run TestIntMin
//go test -bench=.
