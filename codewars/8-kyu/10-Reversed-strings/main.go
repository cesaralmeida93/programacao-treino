package main

import "fmt"

func Solution(word string) (result string) {
	for _, v := range word {
		result = string(v) + result
	}
	return
}


//-------TRAINING-------
// func Solution(word string) (result string) {
// 	for i, v := range word {
// 		fmt.Println("index: ", i)
// 		fmt.Println("value before string: ", v)
// 		fmt.Println("value after string: ", string(v))
// 		fmt.Println("result before: ", result)
// 		result = string(v) + result
// 		fmt.Println("result after: ", result)
// 		fmt.Println("")
// 	}
// 	return result
// }

func main() {
	fmt.Println(Solution("teste"))
}