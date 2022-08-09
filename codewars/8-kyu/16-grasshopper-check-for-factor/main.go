package main

import (
	"fmt"
)

func CheckForFactor(base int, factor int) bool {
  if base % factor == 0 {
    return true
  }
  return false
}

//-------BEST IMPLEMENTATION-------
// package kata

// func CheckForFactor(base int, factor int) bool {
//   return base % factor == 0;
// }

func main() {
	fmt.Println(CheckForFactor(10, 2)) //4
  fmt.Println(CheckForFactor(63, 7)) //4
  fmt.Println(CheckForFactor(2450, 5)) //4
  fmt.Println(CheckForFactor(24612, 3)) //4
  fmt.Println(CheckForFactor(9, 2)) //4
  fmt.Println(CheckForFactor(653, 7)) //4
  fmt.Println(CheckForFactor(2453, 5)) //4
  fmt.Println(CheckForFactor(24617, 3)) //4


}
