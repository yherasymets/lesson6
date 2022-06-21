package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	tokens := strings.Split(input, " ")

	// var min, max string = tokens[0], tokens[0]
	// for _, val := range tokens {
	// 	if min > val {
	// 		min = val
	// 	}
	// 	if max < val {
	// 		max = val
	// 	}
	// }

	sort.Strings(tokens)

	result := tokens[0] + " " + tokens[len(tokens)-1]
	fmt.Println(result)

}
