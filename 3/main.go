package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 ==0 || n%3==0 {
		return false
	}
	i := 5
	for {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
		i+=6
		if i * i > n {
			return true
		}
	}
}

func factorize(n int) ([]int, error) {
	if isPrime(n) {
		return []int{n}, nil
	}


	var factors []int
	original := n
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if !isPrime(i) {
			continue
		}
		// if this prime can divide n w/o leaving any remainder,
		// keep it as a factor of n.
		if n % i == 0 {
			factors = append(factors, i)
			// dive n with i
			n /= i
			// reset the loop
			i = 1
		}
		// if the number we have been left with is a prime then
		// we are done
		if isPrime(n) {
			return append(factors, n), nil
		}
		// if n is zero at this point, we have already
		// found all of its factors
		if n == 0 {
			return factors, nil
		}
	}

	return nil, fmt.Errorf("couldn't factorize %d: found factors %v and remainder %d", original, factors, n)
}


var candidate = flag.Int("n", 591558728, "Run integer factorization for this number")

func main() {
	flag.Parse()

	factors, err := factorize(*candidate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(factors) == 1 && factors[0] == *candidate {
		fmt.Printf("Number %d is prime\n", *candidate)
	} else {
		fmt.Printf("Number %d factorized to %v\n", *candidate, factors)
	}
}
