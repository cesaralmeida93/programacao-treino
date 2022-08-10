package main

import (
	"fmt"
)

func NbDig(n int, d int) int {
	fmt.Println("teste")
	return 1
}

// TRAINING
// import "strings"

// func GetCount(strn string) int {
//   count := 0

//   vowels := []string{"a", "e", "i", "o", "u"}

//   for _, vowel := range vowels {
//     count += strings.Count(strn, vowel)
//   }

//   return count
// }

// BEST SOLUTION
// func GetCount(str string) (count int) {
// 	for _, c := range str {
// 	  switch c {
// 	  case 'a', 'e', 'i', 'o', 'u':
// 		count++
// 	  }
// 	}
// 	return count
//   }

func main() {
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
}
