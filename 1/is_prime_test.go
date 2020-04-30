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
// BenchmarkIsPrime1-4   	     100	  18904583 ns/op
// PASS
// ok  	_/home/michalis/git/comp530/1	1.895s
func BenchmarkIsPrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(rand.Int()%N)
	}
}
