package main

import (
	"flag"
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	divideIfNot := 2
	for {
		if float64(divideIfNot) > math.Sqrt(float64(n)){
			return true
		}
		remainder := n % divideIfNot
		if remainder == 0 {
			return false
		}
		divideIfNot++
	}
}

var candidate = flag.Int("n", 591558727, "Run primality check for this number")

func main() {
	flag.Parse()

	fmt.Println(isPrime(*candidate))
}
