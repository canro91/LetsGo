package roman

import (
	"strings"
)

type Roman struct {
	Value int
	Symbol string
}

var romans = []Roman{
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

func ConvertToRoman(n int) string {
	var builder strings.Builder
	for _, roman := range romans {
		for n >= roman.Value {
			builder.WriteString(roman.Symbol)
			n -= roman.Value
		}
	}
	return builder.String()
}
