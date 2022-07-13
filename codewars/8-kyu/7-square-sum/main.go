package main

import (
	"fmt"
)

func SquareSum(numbers []int) (res int) {
	for i, v := range numbers {
		fmt.Println("key: ", i)
		fmt.Println("Value: ", v)
		fmt.Println("Value in array: ", numbers[i])
		fmt.Println("")
		res += v*v
	}
  return res
}


func main() {
	fmt.Println(SquareSum([]int{1, 2}))
	fmt.Println(SquareSum([]int{0,3,4,5}))
	fmt.Println(SquareSum([]int{}))
}