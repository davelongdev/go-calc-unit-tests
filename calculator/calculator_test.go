package calculator_test

import (
  "testing"
  "github.com/davelongdev/go-calc-unit-tests/calculator"
)

func TestDivide(t *testing.T) {

  // what we expect the function to return
  expected := 2.0

  // what the function actualy returns
  got := calculator.Divide(10.0, 5.0)

  // what to do if got is different than expected
  if got != expected {
    t.Errorf("expected %.1f, got %.1f", expected, got)
  }
}

func TestDivideNegativeDivisor(t *testing.T) {

  // what we expect the function to return
  expected := -2.0

  // what the function actualy returns
  got := calculator.Divide(10.0, -5.0)

  // what to do if got is different than expected
  if got != expected {
    t.Errorf("expected %.1f, got %.1f", expected, got)
  }
}
