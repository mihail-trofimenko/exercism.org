package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, n int) (float64, error) {
	totalAmount, err := fc.FodderAmount(n)
	if err != nil {
		return 0, err
	}

	ff, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return totalAmount / float64(n) * ff, err

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
	err := errors.New("invalid number of cows")
	if n < 1 {
		return 0, err
	}
	return DivideFood(fc, n)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	message string
	details string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%s %s", e.message, e.details)
}

func ValidateNumberOfCows(n int) *InvalidCowsError {
	switch {
	case n < 0:
		return &InvalidCowsError{
			message: fmt.Sprintf("%d cows are invalid:", n),
			details: "there are no negative cows",
		}
	case n == 0:
		return &InvalidCowsError{
			message: fmt.Sprintf("0 cows are invalid:"),
			details: "no cows don't need food",
		}

	}
	return nil

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

/*

package thefarm

// This file contains types used in the exercise and tests and should not be modified.

// FodderCalculator provides helper methods to determine the optimal
// amount of fodder to feed cows.
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}*/
