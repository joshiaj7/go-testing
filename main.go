package main

import (
	"fmt"
)

/*
	To run test from main : go test .
	To run a single test file : go test file.go file_test.go
	To run a single unit test : go test -run TestName
	To see test coverage : go test -cover
*/

func main() {
	fmt.Println("Go testing")

	empty := hello("")
	fmt.Println(empty)

	john := hello("John")
	fmt.Println(john)

	addResult, _ := add(1, 2, 3)
	fmt.Println(addResult)
}
