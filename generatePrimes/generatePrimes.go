package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func isPrime(number int, prevPrimes []int) bool {
	for i := range prevPrimes {
		if number%prevPrimes[i] == 0 {
			return false
		}
	}
	return true
}

func generatePrimes(upper int) (primes []int) {
	for i := 2; i <= upper; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter the upper limit of primes to include as the first argument.")
		os.Exit(2)
	}
	start := time.Now()
	s := os.Args[1]
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(generatePrimes(i))
	elapsed := time.Since(start)
	fmt.Println("Calculated in " + elapsed.String())
}
