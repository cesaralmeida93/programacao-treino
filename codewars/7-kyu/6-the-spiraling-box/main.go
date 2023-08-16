package main

import (
	"fmt"
	"strings"
)

func CreateBox(width, length int) string {
	result := make([][]int, length)
	for i := range result {
		result[i] = make([]int, width)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			distanceToBorder := i
			if j < distanceToBorder {
				distanceToBorder = j
			}
			if width-j-1 < distanceToBorder {
				distanceToBorder = width - j - 1
			}
			if length-i-1 < distanceToBorder {
				distanceToBorder = length - i - 1
			}

			fmt.Printf("i: %d, j: %d, distanceToBorder: %d\n", i, j, distanceToBorder)

			result[i][j] = distanceToBorder + 1

			fmt.Printf("result[%d][%d]: %d\n", i, j, distanceToBorder+1)
		}
	}

	var builder strings.Builder
	for _, row := range result {
		fmt.Fprintln(&builder, row)
	}

	return builder.String()
}

func main() {
	fmt.Println(CreateBox(7, 8))
}

//BEST PRACTICES
// package kata
// import "sort"

// func CreateBox(width, length int) [][]int {
// 	result := [][]int{}
// 	for y := 0; y < length; y++ {
// 		sub := []int{}
// 		for x := 0; x < width; x++ {
// 			numbers := []int{x + 1, y + 1, width - x, length - y}
// 			sort.Ints(numbers)
// 			sub = append(sub, numbers[0])
// 		}
// 		result = append(result, sub)
// 	}
// 	return result
// }

//BEST PRACTICES
// package kata
// func CreateBox(width, length int) [][]int {
// 	result := make([][]int, length)
// 	for y := 0; y < length; y++ {
// 		result[y] = make([]int, width)
// 		for x := 0; x < width; x++ {
// 			result[y][x] = 1 + min(x, y, width-x-1, length-y-1)
// 		}
// 	}
// 	return result
// }

// func min(a, b, c, d int) int {
// 	if a < b {
// 		b = a
// 	}
// 	if c < b {
// 		b = c
// 	}
// 	if d < b {
// 		b = d
// 	}
// 	return b
// }

// package kata
// func CreateBox(width, length int) [][]int {
//     box := make([][]int, length)
//     for i := range box {
//         box[i] = make([]int, width)
//     }

//     for i := 0; i <= length/2; i++ {
//         for j := 0; j <= width/2; j++ {
//             val := 1 + i
//             if j < val {
//                 val = 1 + j
//             }
//             if width-j-1 < val {
//                 val = 1 + width-j-1
//             }
//             if length-i-1 < val {
//                 val = 1 + length-i-1
//             }

//             box[i][j] = val
//             box[length-i-1][j] = val
//             box[i][width-j-1] = val
//             box[length-i-1][width-j-1] = val
//         }
//     }

//     return box
// }
