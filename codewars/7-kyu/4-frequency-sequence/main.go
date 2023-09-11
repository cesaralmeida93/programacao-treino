package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) string {
	frequency := make(map[string]int)
	response := make([]string, len(str))

	fmt.Println("frequency map at beggining: ", len(frequency))
	fmt.Println("response array at begining:", len(response))

	fmt.Println("range str: ", len(str))

	for _, value := range str {
		l := string(value)
		fmt.Printf("%s: %v\n", l, frequency[l])
		fmt.Println("frequency[l]:", frequency[l])
		frequency[l] = frequency[l] + 1
	}

	for index, _ := range response {
		x := string(str[index])
		fmt.Println("x:", x)
		response[index] = strconv.Itoa(frequency[x])
	}

	fmt.Println("frequency: ", frequency)
	fmt.Println("response: ", response)

	fmt.Println("frequency map at end: ", len(frequency))
	fmt.Println("response array at end:", len(response))

	return strings.Join(response, sep) //Code goes here!

}

// TRAINING
// package kata

// import (
// 	"strconv"
// 	"strings"
// )

// func FreqSeq(str string, sep string) string {

// 	buf := map[rune]int{}
// 	result := []string{}

// 	for _, ch := range str {
// 		buf[ch]++
// 	}
// 	for _, ch := range str {
// 		result = append(result, strconv.Itoa(buf[ch]))
// 	}
// 	return strings.Join(result, sep)
// }

func main() {
	fmt.Println(FreqSeq("hello world", "-"))
	// fmt.Println(FreqSeq("19999999", ":"))
	// fmt.Println(FreqSeq("^^^**$", "x"))

}
