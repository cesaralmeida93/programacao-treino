package main

import "fmt"

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

func main() {
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(1))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(7))
	fmt.Println(Factorial(17))
}
