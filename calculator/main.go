package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 2; i++ {
		fmt.Print("Value", i+1, ":")
		input, _ := reader.ReadString('\n')
		float, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println(err)
			panic("Value must be a number")
		}
	}

	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)

}
