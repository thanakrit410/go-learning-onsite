package calculator_test

import (
	"testing"

	"test-go.go/calculator"
)

func TestFizzBuzz(t *testing.T) {
	give := 1
	want := "1"

	got := calculator.Fizzbuzz(give)

	if got != want {
		t.Errorf("given %v wanted %v but got %v", give, got, want)
	}
}
