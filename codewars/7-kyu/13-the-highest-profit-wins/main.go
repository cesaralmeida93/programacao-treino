package main

import "fmt"

func MinMax(arr []int) [2]int {
	if len(arr) == 0 {
		return [2]int{}
	}

	min, max := arr[0], arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return [2]int{min, max}
}

//BEST SOLUTION

// package main
// import "sort"

// func MinMax(arr []int) [2]int {
//   sort.Ints(arr)
//   return [2]int{arr[0], arr[len(arr)-1]}
// }

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinMax([]int{2334454, 5}))
}
