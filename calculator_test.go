package calculator_test

import (
	"testing"

	"github.com/NikolaM-Dev/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	var want float64 = 4
	got := calculator.Add(2, 2)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	var want float64 = 2
	got := calculator.Subtract(4, 2)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	var want float64 = 9
	got := calculator.Multiply(3, 8)

	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
