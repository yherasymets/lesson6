package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}

	result := []int{}

	existedValues := map[int]bool{}

	for _, v := range arr {
		if exists := existedValues[v]; !exists {
			result = append(result, v)
		}
		existedValues[v] = true
	}

	fmt.Println(result)
}

