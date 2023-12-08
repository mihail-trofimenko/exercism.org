/*
Instructions
The ISBN-10 verification process is used to validate book identification numbers.
These normally contain dashes and look like: 3-598-21508-8

ISBN
The ISBN-10 format is 9 digits (0 to 9) plus one check character (either a digit or an X only).
In the case the check character is an X, this represents the value '10'.
These may be communicated with or without hyphens,
and can be checked for their validity by the following formula:

(d₁ * 10 + d₂ * 9 + d₃ * 8 + d₄ * 7 + d₅ * 6 + d₆ * 5 + d₇ * 4 + d₈ * 3 + d₉ * 2 + d₁₀ * 1) mod 11 == 0
If the result is 0, then it is a valid ISBN-10, otherwise it is invalid.

Example
Let's take the ISBN-10 3-598-21508-8. We plug it in to the formula, and get:

(3 * 10 + 5 * 9 + 9 * 8 + 8 * 7 + 2 * 6 + 1 * 5 + 5 * 4 + 0 * 3 + 8 * 2 + 8 * 1) mod 11 == 0
Since the result is 0, this proves that our ISBN is valid.
*/
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func getDigitsSlice(s string) []int {
	if s == "" {
		return nil
	}
	var digits []int
	digits = make([]int, 0)
	var d int
	var lastCharIsX bool

	if s[len(s)-1] == 'X' {
		lastCharIsX = true
	}

	for i, r := range s {
		switch {
		case lastCharIsX && i == len(s)-1:
			digits = append(digits, 10)
			continue
		case r != 'X' && i < len(s)-1 && r != '-' && !unicode.IsDigit(r) && lastCharIsX:
			continue
		case !unicode.IsDigit(r) && i < len(s)-1 && r != '-':
			digits = append(digits, 0)
			continue
		case unicode.IsDigit(r):
			d, _ = strconv.Atoi(string(r))
			digits = append(digits, d)
			continue
		}
	}
	return digits
}

func IsValidISBN(isbn string) bool {
	digits := getDigitsSlice(isbn)
	mul := 1
	sum := 0
	if len(isbn) < 3 || len(digits) > 10 {
		return false
	}

	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * mul
		mul++
	}
	if sum%11 == 0 {
		return true
	}
	return false
}

func main() {
	// Valid ISBN example
	fmt.Println(IsValidISBN("3-598-2X507-9"))
	// Invalid ISBN example
	fmt.Println(IsValidISBN("3-598-2X507-9"))
}
