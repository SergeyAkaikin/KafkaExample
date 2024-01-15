package message

import (
	"strconv"
	"strings"

	ntw "github.com/divan/num2words"
)

type Sync struct {
	Number      int
	Roman       string
	Description string
}

func New(number int) Sync {
	roman := integerToRoman(number)
	description := integerToWord(number)

	return Sync{number, roman, description}
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func integerToWord(number int) string {
	return ntw.Convert(number)
}
