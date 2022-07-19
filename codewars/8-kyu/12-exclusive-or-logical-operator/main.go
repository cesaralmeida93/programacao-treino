package main

import "fmt"

func Xor(a, b bool) bool {
  if a != b {
    return true
  }
  return false
}

//-------TRAINING-------

// func main() {
//   boolA := true
//   boolB := false

//   if boolA != boolB {
//     fmt.Println(true)
//   }

// }

//-------BEST IMPLEMENTATION-------
// package kata

// func Xor(a, b bool) bool {
//   return a != b
// }

func main() {
	fmt.Println(Xor(false, false))
	fmt.Println(Xor(true, false))
	fmt.Println(Xor(false, true))
	fmt.Println(Xor(true, true))
	fmt.Println(Xor(false, Xor(false, false)))
}

