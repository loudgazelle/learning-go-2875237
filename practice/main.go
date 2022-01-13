package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Println(addValues(420, 69))
	sum := addValues(5, 8)
	fmt.Println("The sum is", sum)

	multiSum, multiCount := addAllValues(1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25)
	fmt.Println("Sum of multiple values:", multiSum)
	fmt.Println("count of iteems", multiCount)
}

func doSomething() {
	fmt.Println("doing something")
}

func addValues(value1, value2 int) int {
	return value1 + value2

}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
