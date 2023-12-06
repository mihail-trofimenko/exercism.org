package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {

	// Spaces are allowed in the input, but they should be stripped before checking.
	id = strings.ReplaceAll(id, " ", "")

	// Strings of length 1 or less are not valid.
	if len(id) < 2 {
		return false
	}

	// All other non-digit characters are disallowed.
	for _, r := range id {
		if !unicode.IsDigit(rune(r)) {
			return false
		}
	}
	// The first step of the Luhn algorithm is to double every second digit,
	// starting from the right. We will be doubling
	for i := len(id) - 2; i >= 0; i -= 2 {
		digit, _ := strconv.Atoi(string(id[i]))
		digit *= 2
		// If doubling the number results in a number greater
		// than 9 then subtract 9 from the product.
		if digit > 9 {
			digit -= 9
		}
		id = replaceAtIndex(id, rune('0'+digit), i)
	}
	// Then sum all of the digits
	sum := 0
	for _, r := range id {
		digit, _ := strconv.Atoi(string(r))
		sum += digit
	}
	// If the sum is evenly divisible by 10, then the number is valid. This number is valid!
	if sum%10 == 0 {
		return true
	}
	return false
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func main() {

	id := "4539 3195 0343 6467"
	fmt.Printf("Card ID %s   Vald: %t\n", id, Valid(id))
	id = "8273 1232 7352 0569"
	fmt.Printf("Card ID %s   Vald: %t\n", id, Valid(id))
}
