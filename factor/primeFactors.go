package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func findFactors(number int) (factors []int) {
	// "number / 2" since 2 is the smallet possible prime factor
	i := 2
	for i <= number/2 {
		if number%i == 0 {
			factors = append(factors, findFactors(i)...)
			number = number / i
			continue
		}
		i++
	}
	return append(factors, number)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter an integer to factor as the first argument.")
		os.Exit(2)
	}
	start := time.Now()
	s := os.Args[1]
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(findFactors(i))
	elapsed := time.Since(start)
	fmt.Println("Calculated in " + elapsed.String())
}
