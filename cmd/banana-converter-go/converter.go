package main

import "fmt"

const (
	BananaToInch      = 7.0
	BananaToCm        = 17.78
	BananaToFeet      = 0.583
	BananaToMeter     = 0.178
	BananaToYard      = 0.194
	BananaToKilometer = 0.000178
	BananaToMile      = 0.00011
)

var errFoo = fmt.Errorf("invalid unit conversion")

func convert(value float64, fromUnit string, toUnit string) (float64, error) {
	conversion := map[string]float64{
		"banana":    1,
		"inch":      BananaToInch,
		"cm":        BananaToCm,
		"feet":      BananaToFeet,
		"meter":     BananaToMeter,
		"yard":      BananaToYard,
		"kilometer": BananaToKilometer,
		"mile":      BananaToMile,
	}
	fromConversion, fromExists := conversion[fromUnit]
	toConversion, toExists := conversion[toUnit]
	if !fromExists || !toExists {
		return 0, errFoo
	}

	lengthValue := value / fromConversion
	result := lengthValue * toConversion

	return result, nil
}
