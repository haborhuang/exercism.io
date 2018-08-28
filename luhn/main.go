package main

import (
	"log"
	"strings"
	"unicode"
)

func main() {
	var testCases = []struct {
		description string
		input       string
		ok          bool
	}{
		{
			"single digit strings can not be valid",
			"1",
			false,
		},
		{
			"a single zero is invalid",
			"0",
			false,
		},
		{
			"a simple valid SIN that remains valid if reversed",
			"059",
			true,
		},
		{
			"a simple valid SIN that becomes invalid if reversed",
			"59",
			true,
		},
		{
			"a valid Canadian SIN",
			"055 444 285",
			true,
		},
		{
			"invalid Canadian SIN",
			"055 444 286",
			false,
		},
		{
			"invalid credit card",
			"8273 1232 7352 0569",
			false,
		},
		{
			"valid strings with a non-digit included become invalid",
			"055a 444 285",
			false,
		},
		{
			"valid strings with punctuation included become invalid",
			"055-444-285",
			false,
		},
		{
			"valid strings with symbols included become invalid",
			"055£ 444$ 285",
			false,
		},
		{
			"single zero with space is invalid",
			" 0",
			false,
		},
		{
			"more than a single zero is valid",
			"0000 0",
			true,
		},
		{
			"input digit 9 is correctly converted to output digit 9",
			"091",
			true,
		},
		{
			"strings with non-digits is invalid",
			":9",
			false,
		},
	}

	for _, c := range testCases {
		if Luhn(c.input) != c.ok {
			log.Fatalf("Case '%s' fail", c.description)
		}
	}
}

func Luhn(str string) bool {
	str = strings.TrimSpace(str)

	i := 0
	sum := 0

	for _, r := range str {
		b := byte(r)

		// Skip blank
		if b == ' ' {
			continue
		}

		// There is non-digit char
		if !unicode.IsDigit(r) {
			return false
		}

		// Sum up
		n := int(b - '0')
		if (len(str)-i)%2 == 0 {
			// Double every second digit, starting from the right
			sum += 2 * n
			if n > 4 {
				sum -= 9
			}
		} else {
			sum += n
		}

		i++
	}

	return sum%10 == 0 && (sum != 0 || len(str) > 1)
}
