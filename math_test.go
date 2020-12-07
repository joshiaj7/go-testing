package main

import (
	"testing"
)

type TestAddInput struct {
	Data   []int
	Result int
}

func TestAdd(t *testing.T) {
	inputs := []TestAddInput{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-1, -2, -3}, -6},
		{[]int{}, 0},
	}

	for _, item := range inputs {
		result, err := add(item.Data...)
		if err != nil {
			t.Errorf("%v", err)
		}

		if result == item.Result {
			t.Logf("Add() func success")
		} else {
			t.Errorf("Add() func failed, expected result %v, got %v", item.Result, result)
		}
	}
}
