package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	reversedString := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversedString += string(str[i])
	}

	if strings.ToLower(str) == strings.ToLower(reversedString) {
		return true
	}

	return false
}

// best soluction
// package main

// import "strings"

// func IsPalindrome(str string) bool {
//   str = strings.ToLower(str)
//   n := len(str)
//   for i := 0; i < n; i++ {
//     n -= 1
//     if str[i] != str[n] {
//       return false
//     }
//   }
//   return true
// }
func main() {
	fmt.Println(IsPalindrome("a"))
	fmt.Println(IsPalindrome("aba"))
	fmt.Println(IsPalindrome("Abba"))
	fmt.Println(IsPalindrome("hello"))
}
