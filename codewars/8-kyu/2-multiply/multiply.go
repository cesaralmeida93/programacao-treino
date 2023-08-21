package main

import "fmt"

func Multiply(a, b int) int {
	return a * b
}

// package main

// import "fmt"

// type fattori struct {
// 	fattore1 int
// 	fattore2 int
// }

// func Multiply(a, b int) int {

// 	moltiplicazione := fattori{fattore1: a,
// 		fattore2: b}

// 	return moltiplicazione.fattore1 * moltiplicazione.fattore2
// }

func main() {
	fmt.Println(Multiply(3, 3))
}
