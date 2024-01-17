package main

import (
	"errors"
	"fmt"
)

var conversions = map[string]float64{
	"banana":    1,
	"inch":      7.0,
	"cm":        17.78,
	"feet":      0.583,
	"meter":     0.178,
	"yard":      0.194,
	"kilometer": 0.000178,
	"mile":      0.00011,
}

func convert(value float64, fromUnit string, toUnit string) (float64, error) {
	fromConversion, fromExists := conversions[fromUnit]
	toConversion, toExists := conversions[toUnit]

	if !fromExists || !toExists {
		return 0, errors.New("Invalid unit")
	}

	lengthInBananas := value / fromConversion
	result := lengthInBananas * toConversion

	return result, nil
}

func main() {
	value := 10.0
	fromUnit := "banana"
	toUnit := "cm"

	result, err := convert(value, fromUnit, toUnit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%v %s is equal to %v %s\n", value, fromUnit, result, toUnit)
	}
}
