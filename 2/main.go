package main

import (
	"flag"
	"fmt"
)

// NOTE: I couldn't think of any better way to find
// a prime than my first solution so I searched online.
// According to Wikipedia[1], "all integers can be expressed
// as (6k + i) for some integer k and for i = −1, 0, 1, 2, 3, or 4;
// 2 divides (6k + 0), (6k + 2), (6k + 4); and 3 divides (6k + 3).
// So, a more efficient method is to test if n is divisible by 2 or
// 3, then to check through all the numbers of the form 6k±1 ≤ n.
// This is 3 times as fast as testing all m."
//
// Indeed, in my benchmarks, this method is at least 4-5 times faster
// than the isPrime method in the first solution.
//
// [1] https://en.wikipedia.org/wiki/Primality_test
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


var candidate = flag.Int("n", 5915587277, "Run primality check for this number")

func main() {
	flag.Parse()

	fmt.Println(isPrime(*candidate))
}
