package main

import (
	"fmt"
	"sort"
)

func Gimme(array [3]int) int {
	fmt.Printf("array: %v\n", array)
	arrayCopy := array
	arrSlice := arrayCopy[:]
	sort.Ints(arrSlice)
	val := arrSlice[1]

	for i, v := range array {
		if v == val {
			return i
		}
	}

	return 0
}

func main() {
	fmt.Println(Gimme([3]int{2, 3, 1}))
	fmt.Println(Gimme([3]int{5, 10, 14}))
}
