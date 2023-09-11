package main

import (
	"fmt"
)

func Add(i int) func(int) int {
	return func(a int) int {
		return a + i
	}
}

func main() {
	fmt.Println(Add(1)(3))
}
