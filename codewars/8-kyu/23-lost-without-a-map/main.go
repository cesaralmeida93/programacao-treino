package main

import "fmt"

func Maps(x []int) []int {
	result := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		result[i] = x[i] * 2
	}
	return result
}
func main() {
	fmt.Println(Maps([]int{1, 2, 3}))
	fmt.Println(Maps([]int{4, 1, 1, 1, 4}))
	fmt.Println(Maps([]int{2, 2, 2, 2, 2, 2}))
}
