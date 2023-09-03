package main

import "fmt"

func ReverseSeq(n int) []int {
	result := make([]int, 0, n)

	for i := n; i >= 1; i-- {
		result = append(result, i)
	}

	return result
}

// best solution
// package main

// func ReverseSeq(n int) []int {
//   arr := make([]int, n)
//   for i:= 0; i < n; i++ {
//     arr[i] = n - i
//   }
//   return arr
// }
func main() {
	fmt.Println(ReverseSeq(5))
}
