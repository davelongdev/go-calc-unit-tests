package calculator_test

import (
  "testing"
  "errors"
  "github.com/davelongdev/go-calc-unit-tests/calculator"
)


// create struct for test cases
var testcases = []struct {
  name string
  dividend float64
  divisor float64
  expected float64
  expectedError error
}{
  // add test cases
  {"division", 10.0, 5.0, 2.0, nil},
  {"division by negative value", 10.0, -5.0, -2.0, nil},
  {"division by zero", 2.0, 0.0, 0.0, errors.New("division by zero not permitted")},
}

func TestDivide(t *testing.T) {

  // loop over the test cases
  for _, tc := range testcases {
    t.Run(tc.name, func(t *testing.T) {

      // what the function is expected to return
      expected := tc.expected

      // what the function actually returns
      gotValue, gotError := calculator.Divide(tc.dividend, tc.divisor)
      
      // check if divide function returns an error
      if gotError != nil {

        // compare error strings of both errors
        if gotError.Error() != tc.expectedError.Error() {
          t.Errorf("expected error %s, got %s", tc.expectedError.Error(), gotError.Error())
        }
      }

      // what to do if got is different than expected
      if gotValue != expected {
        t.Errorf("expected %.1f, got %.1f", expected, gotValue)
      }
    })
  }
}
