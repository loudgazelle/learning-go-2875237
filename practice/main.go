package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "it's Sunday!"
		// fallthrough
	case 2:
		result = "it's Monday!"
		// fallthrough
	case 3:
		result = "it's Tuesday!"
		// fallthrough
	case 4:
		result = "it's Wednesday!"
		// fallthrough
	default:
		result = "it's some other day!"
	}
	fmt.Println(result)
}
