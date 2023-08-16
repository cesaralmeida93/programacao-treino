package main

import (
	"fmt"
)

func RemoveChar(word string) string {
	wordWithoutFirstAndLast := word[1 : len(word)-1]

	return string(wordWithoutFirstAndLast)

}

func main() {
	fmt.Println(RemoveChar("eloquent"))
	fmt.Println(RemoveChar("country"))
	fmt.Println(RemoveChar("person"))
	fmt.Println(RemoveChar("place"))

}
