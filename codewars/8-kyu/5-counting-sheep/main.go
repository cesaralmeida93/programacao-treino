package main

import "fmt"

func CountSheeps(numbers []bool) int {
	var res int

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == true {
			res += 1
		}
	}

  return res
}

//-------BEST PRACTICE-------
// package kata

// func CountSheeps(numbers []bool) (cnt int) {
//     for _, b := range numbers {
//         if b {
//             cnt++;
//         }
//     }
//     return cnt
// }

func main() {
	arr1 := []bool{
		true,  true,  true,  false,
		true,  true,  true,  true ,
		true,  false, true,  false,
		true,  false, false, true ,
		true,  true,  true,  true ,
		false, false, true,  true,
	}
	fmt.Println(CountSheeps(arr1))
}