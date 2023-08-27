// package main

// import (
// 	"fmt"
// )

// func Dominator(a []int) int {
// 	// indexCount := make(map[int]int)
// 	valueCount := make(map[int]int)
// 	// fmt.Println("indexCount outside for: ", indexCount)
// 	fmt.Println("valueCount outside for: ", valueCount)

// 	for _, value := range a {
// 		// indexCount[index]++
// 		valueCount[value]++
// 		// fmt.Println("indexCount inside for: ", indexCount)
// 		fmt.Println("valueCount inside for: ", valueCount)
// 		fmt.Printf("valueCount[%d]: %d\n", value, valueCount[value])
// 		fmt.Println("value: ", value)

// 		if valueCount[value] > len(a)/2 {
// 			return value
// 		}
// 	}

// 	return -1
// }

// func main() {
// 	for _, list := range [][]int{
// 		{3, 4, 3, 2, 3, 1, 3, 3},
// 		{1, 2, 3, 4, 5},
// 		{3, 4, 3, 2, 3, 1, 3, 3},
// 		{1, 1, 1, 2, 2, 2},
// 		{1, 1, 1, 2, 2, 2, 2},
// 	} {
// 		fmt.Printf("%v => %d\n", list, Dominator(list))
// 	}
// }

package main

import "fmt"

func dominator(arr []int) int {
	numCount := make(map[int]int)

	for _, num := range arr {
		numCount[num]++

		if numCount[num] > len(arr)/2 {
			return num
		}
	}

	return -1
}

func main() {
	for _, list := range [][]int{
		{3, 4, 3, 2, 3, 1, 3, 3},
		{1, 2, 3, 4, 5},
		{3, 4, 3, 2, 3, 1, 3, 3},
		{1, 1, 1, 2, 2, 2},
		{1, 1, 1, 2, 2, 2, 2},
	} {
		fmt.Printf("%v => %d\n", list, dominator(list))
	}
}
