package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' functions
func DivideFood(calc FodderCalculator, numCows int) (float64, error) {
	amt, err := calc.FodderAmount(numCows)
	if err != nil {
		return 0.0, err
	}
	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return amt / float64(numCows) * factor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(calc, numCows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	numCows       int
	customMessage string
}

func (customErr *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", customErr.numCows, customErr.customMessage)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{numCows, "there are no negative cows"}
	} else if numCows == 0 {
		return &InvalidCowsError{numCows, "no cows don't need food"}
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
