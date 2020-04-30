package main

import (
	"math/rand"
	"testing"
)

// limiting the numbers up to 15 digit to keep
// the benchmarks fast.
const N = 1e15

// Example run:
//
// $ make profile
// goos: linux
// goarch: amd64
// BenchmarkIsPrime2-4   	    1873	    669441 ns/op
// PASS
// ok  	_/home/michalis/git/comp530/2	3.072s
func BenchmarkIsPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(rand.Int()%N)
	}
}
