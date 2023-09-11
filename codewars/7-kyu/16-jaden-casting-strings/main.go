package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) (res string) {
	for i, c := range str {
		if i == 0 || str[i-1] == ' ' {
			res += strings.ToUpper(string(c))
		} else {
			res += string(c)
		}
	}
	return
}

func main() {
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
}
