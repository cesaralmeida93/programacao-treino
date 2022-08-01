package main

import (
	"fmt"
)

func Multiply(a int, b int) int {
  return a * b
}

//-------BEST IMPLEMENTATION-------
// package kata

// func Multiply(a int, b int) int {
//   return a * b
// }

func main() {
	fmt.Println(Multiply(2, 2)) //4
  fmt.Println(Multiply(4, 5)) //20
  fmt.Println(Multiply(0, 100)) //0
  fmt.Println(Multiply(1, 100)) //100


}
