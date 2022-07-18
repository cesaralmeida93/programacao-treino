package main

import "fmt"

func ToAlternatingCase(str string) string {
  rs := make([]rune, len(str))
	fmt.Println("rs: ", rs)
  for i, c := range str {
    r := c
    if c >= 'a' && c <= 'z' {
      r = c + 'A' - 'a'
    }
    if c >= 'A' && c <= 'Z' {
      r = c + 'a' - 'A'
    }
    rs[i] = r
  }
  return string(rs)
}

// -------TRAINING-------
// func main() {
// 	stringi := "HELLO WORLD"
// 	fmt.Println("string: ", stringi)
// 	stringiArray := make([]rune, len(stringi)) 
// 	fmt.Println("stringiArray: ", stringiArray)
// 	for key, value := range stringi {
// 		resultKey := key
// 		fmt.Println("key: ", resultKey)
// 		resultValue := value
// 		fmt.Println("value: ", resultValue)
// 		if value >= 'a' && value <= 'z' {
// 			resultValue = value + 'A' - 'a'
// 			fmt.Println("resultValue: ", resultValue)
// 		} else if value >= 'A' && value <= 'Z' {
// 			resultValue = value + 'a' - 'A'
// 			fmt.Println("resultValue: ", resultValue)
// 			fmt.Println("")
// 		}
// 		stringiArray[key] = resultValue
// 		fmt.Println("stringiArray at position ", key, ":", string(resultValue))
// 	}

// 	fmt.Println("final result: ", string(stringiArray))

// }


//-------BEST IMPLEMENTATION-------
// package kata

// import "unicode"

// func ToAlternatingCase(str string) string {
//   result := []rune{}
//   for _, ch:=range str{
//     if unicode.IsUpper(ch){
//       result = append(result, unicode.ToLower(ch))
//     } else if unicode.IsLower(ch){
//        result = append(result, unicode.ToUpper(ch))
//     } else {
//       result = append(result, ch)
//     }
//   }
  
//   return string(result)
// }

func main() {
	fmt.Println(ToAlternatingCase("hello world"))
	fmt.Println(ToAlternatingCase("HELLO WORLD"))
	fmt.Println(ToAlternatingCase("hello WORLD"))
	fmt.Println(ToAlternatingCase("HeLLo WoRLD"))
	fmt.Println(ToAlternatingCase("12345"))
	fmt.Println(ToAlternatingCase("1a2b3c4d5e"))
	fmt.Println(ToAlternatingCase("String.prototype.toAlternatingCase"))
}