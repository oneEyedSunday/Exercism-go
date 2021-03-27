package romannumerals

import (
	"errors"
	"strings"
)

type arabicRoman struct {
	value int
	digit string
}

var conversions  = [...]arabicRoman{
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

// ToRomanNumeral cobverts from Arabic to Roman Numerals
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 3001 {
		return "", errors.New("invalid input")
	}
	var roman strings.Builder
	for _, conversion := range conversions {
		for input >= conversion.value {
			roman.WriteString(conversion.digit)
			input -= conversion.value
		}
	}
	return roman.String(), nil
}
