package main

import (
	"fmt"
)

func GetSize(w, h, d int) [2]int {
	return [2]int{(2 * d * w) + (2 * d * h) + (2 * h * w), w * h * d}
	// your code here
}

func main() {
	fmt.Println(GetSize(4, 2, 6))
	fmt.Println(GetSize(1, 1, 1))
	fmt.Println(GetSize(1, 2, 1))
	fmt.Println(GetSize(1, 2, 2))
	fmt.Println(GetSize(10, 10, 10))

}
