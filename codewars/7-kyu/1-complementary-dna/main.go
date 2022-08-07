package main

import (
	"fmt"
)

func DNAStrand(dna string) string {
  var dict = map[string]string{
		"A": "T", "T": "A", "C": "G", "G": "C",
	}

	res := ""

	for _, v := range dna {
		res += dict[string(v)]
	}

	return res
}

// TRAINING
// func DNAStrand(dna string) string {
//   var dict = map[string]string{
// 		"A": "T", "T": "A", "C": "G", "G": "C",
// 	}
// 	fmt.Println("dict: ", dict)

// 	for k, v := range dict {
// 		fmt.Println("position: ", k)
// 		fmt.Println("value: ", v)
// 	}

// 	fmt.Println("dict[A]: ", dict[0])
// 	fmt.Println("dict[T]: ", dict[1])
// 	fmt.Println("dict[C]: ", dict[2])
// 	fmt.Println("dict[G]: ", dict[3])
// 	res := ""

// 	for _, v := range dna {
// 		fmt.Println("dict[v]: ", dict[string(v)])
// 		res += res + dict[string(v)]
// 		fmt.Println("res: ", res)
// 	}

// 	return res
// }


// BEST SOLUTION
// package main

// import (
// 	"fmt"
// 	"strings"
// )


// func DNAStrand(dna string) string {
//   return strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G").Replace(dna)
// }

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}