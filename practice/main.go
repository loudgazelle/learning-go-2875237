package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 9)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 69
	numbers[3] = 420
	numbers[4] = 666
	numbers[5] = 484
	numbers[6] = 48
	numbers[7] = 14
	numbers[8] = 34
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
