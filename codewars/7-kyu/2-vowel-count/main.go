package main

import (
	"fmt"
)

func GetCount(str string) (count int) {
	vows := []string{"a", "e", "i", "o", "u"}

	for _, v := range str {
		for _, j := range vows {
			if string(v) == j {
				count++
			}

		}
	}
	// Enter solution here
	return count
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
	fmt.Println(GetCount("abracadabra"))
}
