package main

import (
	"fmt"
	"sort"
)

func ExpressionMatter(a int, b int, c int) int {
  arr := []int{a*(b+c), a*b*c, a+b+c, (a+b)*c}
  sort.Ints(arr)
  return arr[len(arr)-1]
}

//-------BEST IMPLEMENTATION-------
// package kata

// import "sort"

// func ExpressionMatter(a int, b int, c int) int {
//     arr := []int{a*(b+c), a*b*c, a+b+c, (a+b)*c}
//     sort.Ints(arr)
//     return arr[len(arr)-1]
// }

func main() {
	fmt.Println(ExpressionMatter(2, 1, 2)) //6
	fmt.Println(ExpressionMatter(2, 1, 1)) //4
	fmt.Println(ExpressionMatter(1, 1, 1)) //3
	fmt.Println(ExpressionMatter(1, 2, 3)) //9
	fmt.Println(ExpressionMatter(1, 3, 1)) //5
  fmt.Println(ExpressionMatter(2, 2, 2)) //8
  fmt.Println(ExpressionMatter(5, 1, 3)) //20
  fmt.Println(ExpressionMatter(3, 5, 7)) //105
  fmt.Println(ExpressionMatter(5, 6, 1)) //35
  fmt.Println(ExpressionMatter(1, 6, 1)) //8
  fmt.Println(ExpressionMatter(2, 6, 1)) //14
  fmt.Println(ExpressionMatter(6, 7, 1)) //48
  fmt.Println(ExpressionMatter(2, 10, 3)) //60
  fmt.Println(ExpressionMatter(1, 8, 3)) //27
  fmt.Println(ExpressionMatter(9, 7, 2)) //126
  fmt.Println(ExpressionMatter(1, 1, 10)) //20
  fmt.Println(ExpressionMatter(9, 1, 1)) //18
  fmt.Println(ExpressionMatter(10, 5, 6)) //300
  fmt.Println(ExpressionMatter(1, 10, 1)) //12

}

