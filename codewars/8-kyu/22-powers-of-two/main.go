package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) (arr []uint64) {
	var pot, i float64
	for i = 0; i <= float64(n); i++ {
		pot = math.Pow(2, i)
		arr = append(arr, uint64(pot))
	}
	return arr
}
func main() {
	fmt.Println(PowersOfTwo(0))
	fmt.Println(PowersOfTwo(1))
	fmt.Println(PowersOfTwo(4))
}
