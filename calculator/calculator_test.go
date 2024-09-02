package calculator_test

import (
  "testing"
  "github.com/davelongdev/go-calc-unit-tests/calculator"
)


// create struct for test cases
var testcases = []struct {
  name string
  dividend float64
  divisor float64
  expected float64
}{
  // add test cases
  {"division", 10.0, 5.0, 2.0},
  {"division by negative value", 10.0, -5.0, -2.0},
}

func TestDivide(t *testing.T) {

  // loop over the test cases
  for _, tc := range testcases {
    t.Run(tc.name, func(t *testing.T) {
      // what we expect the function to return
      expected := tc.expected

      // what the function actualy returns
      got := calculator.Divide(tc.dividend, tc.divisor)

      // what to do if got is different than expected
      if got != expected {
        t.Errorf("expected %.1f, got %.1f", expected, got)
      }
    })
  }
}
