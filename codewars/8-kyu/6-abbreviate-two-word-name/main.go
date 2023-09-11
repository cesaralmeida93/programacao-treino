package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	fmt.Println(name)
	words := strings.Split(name, " ")
	fmt.Println(words)
	fmt.Println(len(words))
	fmt.Println(string(words[0]))
	fmt.Println(string(words[0][0]))
	res := ""
	res += string(words[0][0])
	fmt.Println(string(words[0][0]))
	res += "."
	res += string(words[1][0])
	fmt.Println(string(words[1][0]))
	return strings.ToUpper(res)
}

//-------TRAINING-------
// func AbbrevName(name string) string{
// 	words := strings.Split(name, " ")
// 	fmt.Println(string(words[0][0]))
// 	res := ""
// 	res += string(words[0][0])
// 	fmt.Println(string(words[0][0]))
// 	res += "."
// 	res += string(words[1][0])
// 	fmt.Println(string(words[1][0]))
//   return strings.ToUpper(res)
// }

//-------BEST PRACTICE-------
// package kata

// import "strings"

// func AbbrevName(name string) string{
//   words := strings.Split(name, " ")
//   return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
// }

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
}
