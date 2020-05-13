package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GCD returns the greatest common divisor via the euclidean algorithm.
func GCD(num1 int, num2 int) int {
	// makes sure num2 >= num1
	if num1 > num2 {
		num2, num1 = num1, num2
	}
	r := num2 % num1
	if r == 0 {
		return num1
	}
	return GCD(r, num1)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter two integers as arguments")
		os.Exit(2)
	}
	start := time.Now()
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(2)
	}
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(2)
	}

	fmt.Println(GCD(num1, num2))
	elapsed := time.Since(start)
	fmt.Println("Calculated in " + elapsed.String())
}