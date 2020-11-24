package main

import (
	"fmt"
)

func main(){
	fmt.Println("Go testing")

	empty := hello("")
	fmt.Println(empty)

	john := hello("John")
	fmt.Println(john)

	addResult, _ := add(1,2,3)
	fmt.Println(addResult)
}