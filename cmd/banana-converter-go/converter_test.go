// converter_test.go

package main

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		value    float64
		fromUnit string
		toUnit   string
		expected float64
		err      error
	}{
		{1.0, "banana", "inch", 7.0, nil},
		{7.0, "inch", "banana", 1, nil},
		{1.5, "banana", "cm", 26.67, nil},
		{17.78, "cm", "banana", 1, nil},
		{1.0, "banana", "feet", 0.583, nil},
		{0.583, "feet", "banana", 1, nil},
		{1.0, "banana", "meter", 0.178, nil},
		{215.5, "meter", "banana", 1210.67, nil},
		{1.0, "banana", "yard", 0.194, nil},
		{0.194, "yard", "banana", 1, nil},
		{1.0, "banana", "kilometer", 0.000178, nil},
		{0.000178, "kilometer", "banana", 1, nil},
		{1.0, "banana", "mile", 0.00011, nil},
		{0.00011, "mile", "banana", 1, nil},
		{0, "invalid", "cm", 0, errFoo},
	}
	for _, test := range tests {
		result, err := convert(test.value, test.fromUnit, test.toUnit)

		// Format the expected and actual values with 2 decimal places
		expectedFormatted := fmt.Sprintf("%.2f", test.expected)
		resultFormatted := fmt.Sprintf("%.2f", result)

		// Check the result
		if expectedFormatted != resultFormatted {
			t.Errorf("Expected %v %s to be equal to %v %s, but got %v %s", test.value, test.fromUnit, expectedFormatted, test.toUnit, resultFormatted, test.toUnit)
		} // Check the error
		if (err == nil && test.err != nil) || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf("Expected error: %v, but got error: %v", test.err, err)
		}
	}
}
