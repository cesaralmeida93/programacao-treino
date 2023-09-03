package main

import (
	"strings"
)

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	// Implement me! :)
	var TupleList []Tuple = []Tuple{}
	var hashMap = make(map[rune]int)
	for _, c := range text {
		if _, ok := hashMap[c]; !ok {
			v := strings.Count(text, string(c))
			hashMap[c] = v
			TupleList = append(TupleList, Tuple{Char: c, Count: v})
		}
	}

	return TupleList
}

// best implementation

// import "strings"

// type Tuple struct {
// 	Char  rune
// 	Count int
// }

// func OrderedCount(text string) []Tuple {
//   counts := make([]Tuple, 0, len(text))
//   counted := ""

//   for _, r := range(text) {
//     if strings.Count(counted, string(r)) == 0 {
//       counts = append(counts, Tuple{Char: r, Count: strings.Count(text, string(r))})
//       counted += string(r)
//     }
//   }
//   return counts
// }

// func main() {
// 	fmt.Println(OrderedCount("abracadabra"))
// }
