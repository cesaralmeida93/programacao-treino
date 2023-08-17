package main

func EachCons(arr []int, n int) [][]int {
	// your code here
	var result [][]int
	length := len(arr)
	if n > length {
		return result
	}

	for i := 0; i <= length-n; i++ {
		subset := make([]int, n)
		copy(subset, arr[i:i+n])
		result = append(result, subset)
	}

	return result
}

func main() {

}
