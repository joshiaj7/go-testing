package main

import (
	"errors"
)

// add function returns sum of all numbers
func add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) == 0 {
		errorMessage := "There is no number to add"
		return sum, errors.New(errorMessage)
	}

	for _, v := range numbers {
		sum += v
	}

	return sum, nil
}
