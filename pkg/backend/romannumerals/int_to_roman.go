package romannumerals

import (
	"strconv"
	"strings"
)

type RomanIntConversion struct {
	value int
	digit string
}

var conversions = []RomanIntConversion{
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

func IntegerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
