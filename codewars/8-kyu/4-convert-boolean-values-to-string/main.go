package main

import "fmt"

func BoolToWord(word bool) string {
	if word == true {
		return "Yes"
	} else {
		return "No"
	}
}

//-------BEST PRACTICE-------
// package kata

// func BoolToWord(word bool) string {
//     if word {
//       return "Yes"
//     }
    
//     return "No"
// }

func main() {
	fmt.Println(BoolToWord(true))
	fmt.Println(BoolToWord(false))
}