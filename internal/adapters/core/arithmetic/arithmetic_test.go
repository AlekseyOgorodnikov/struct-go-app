package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubstraction(t *testing.T) {
	arith := NewAdapter()
	anwer, err := arith.Substraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, anwer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()
	anwer, err := arith.Multiplication(2, 3)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, anwer, int32(6))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()
	anwer, err := arith.Devision(6, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, anwer, int32(3))
}
