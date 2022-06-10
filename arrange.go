package main

import "fmt"

func main() {

	fmt.Println("No. of digits")
	var num int
	fmt.Scanln(&num)

	input := make([]int, num)

	for i := 0; i < num; i++ {
		fmt.Println("enter digits")
		var digits int
		fmt.Scanln(&digits)
		input[i] = digits
	}
	fmt.Println("input: ", input)
	output := processArray(input)
	fmt.Println("output: ", output)
}

func processArray(arr []int) []int {
	result := []int{}
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 1 {
			result = append(result, arr[i])
		} else {
			count++
		}
	}
	for i := 0; i < count; i++ {
		result = append(result, 1)

	}
	return result
}
