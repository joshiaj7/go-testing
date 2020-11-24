package main

import (
	"fmt"
)


// hello will return hello world
func hello(s string) string {
	if len(s) == 0 {
		return "Hello Dude!"
	}
	return fmt.Sprintf("Hello %s!", s)
}
