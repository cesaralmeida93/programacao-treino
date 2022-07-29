package main

import (
	"fmt"
)

func MakeNegative(x int) int {
  if x >= 0 {
    return x * (-1)
  }
  return x * 1
}

//-------BEST IMPLEMENTATION-------
// package kata

// func MakeNegative(x int) int {

//   if x > 0 {
//     return -x
//   }
//   return x

// }

func main() {
	fmt.Println(MakeNegative(42)) //-42


}

