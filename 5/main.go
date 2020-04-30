package main

import (
	"flag"
	"fmt"
	"os"
)

// Euclid's algorithm can be used to find whether two numbers
// a and m are coprime, ie, gcd(a,m)=1. The Euclidean algorithm
// can also be extended to compute the coefficients of Bézout's
// identity, which are integers x and y such that:
//
// ax + my = gcd(a,m)
//
// The multiplicative inverse of a modulo m exists if and only if
// a and m are coprime. If the modular multiplicative inverse of a
// modulo m exists, the operation of division by a modulo m can be
// defined as multiplying by the inverse.
//
// ax + my = 1
//
// Applying the modulo operation on both sides reduces the above
// equation to:
//
// ax + my mod m = 1 mod m =>
// ax + 0y = 1 (mod m) =>
// ax = 1 (mod m)
//
// Thus, x is the modular multiplicative inverse of a modulo m.
func inverse(a, m int) (int, error) {
	// Zero has no modular multiplicative inverse.
	if a == 0 {
		return 0, nil
	}
	t, newT := 0, 1
	r, newR := m, a
	for newR != 0 {
		q := r / newR
		t, newT = newT, t - q * newT
		r, newR = newR, r - q * newR
	}
	// The multiplicative inverse of a modulo m exists if and only
	// if a and m are coprime.
	if r > 1 {
		return 0, fmt.Errorf("%d is not invertible", a)
	}
	// For getting a result which is positive and lower than m, one may
	// use the fact that the integer t provided by the algorithm satisfies
	// |t| < m. That is, if t < 0, one must add m to it at the end.
	if t < 0 {
		t+=m
	}
	return t, nil
}

var (
	toInvert = flag.Int("x", 342952340, "Number to invert")
	modulo = flag.Int("m", 4230493243, "Modulo over which to invert numbers")
)


func main() {
	flag.Parse()

	inv, err := inverse(*toInvert, *modulo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(inv)
}
