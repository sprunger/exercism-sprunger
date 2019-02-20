package romannumerals

import (
	"fmt"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

var valueMapping = [...]arabicToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral returns string representation of integer [1..3000]
func ToRomanNumeral(input int) (result string, err error) {
	if input < 1 || input > 3000 {
		return result, fmt.Errorf("The value %d is out of range [1..3000]", input)
	}

	for _, value := range valueMapping {
		for input >= value.arabic {
			result += value.roman
			input -= value.arabic
		}
	}

	return result, err
}
