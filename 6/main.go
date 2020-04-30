package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
)

func decrypt(ciphertext, privateKey, modulo int) int {
	cipherBig := big.NewInt(int64(ciphertext))
	privBig := big.NewInt(int64(privateKey))
	modBig := big.NewInt(int64(modulo))
	plaintext := new(big.Int).Exp(cipherBig, privBig, modBig)
	return int(plaintext.Int64())
}

func encrypt(plaintext, publicKey, modulo int) int {
	plainBig := big.NewInt(int64(plaintext))
	pubBig := big.NewInt(int64(publicKey))
	modBig := big.NewInt(int64(modulo))
	ciphertext := new(big.Int).Exp(plainBig, pubBig, modBig)
	return int(ciphertext.Int64())}

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

func gcd(a, b int) int {
	x, y := a, b

	for y != 0 {
		r := x % y
		x = y
		y = r
	}

	return x
}

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


func phi(n, e int) (int, error) {
	factors, err := factorize(n)
	if err != nil {
		return 0, err
	}
	if len(factors) != 2 {
		return 0, fmt.Errorf("need n to be composed of two primes")
	}
	phi := (factors[0]-1)*(factors[1]-1)
	if gcd(phi, e) != 1 {
		return 0, fmt.Errorf("need φ(n) and e to be co-prime numbers")
	}
	return phi, nil
}

var (
	n = flag.Int("n", 937513, "The result of the pq multiplication, used as the modulo operation")
	e = flag.Int("pub", 638471, "Public key")
	plaintext = flag.Int("plaintext", 123456, "Plaintext to encrypt")
)

func main(){
	if *plaintext < 1 || *plaintext >= *n {
		fmt.Printf("Plaintext (%d) must be between 1 and %d\n", *plaintext, *n-1)
		os.Exit(1)
	}
	// Given n, we need to find the primes p and q in order to compute φ(n)
	// such that gcd(φ(n),e)=1.
	phiN, err := phi(*n, *e)
	if err != nil {
		fmt.Printf("cannot get φ(n) as a coprime of e: %v\n", err)
		os.Exit(1)
	}
	// Now that we have φ(n), we can find the private key for e by calculating
	// the modular multiplicative inverse of e modulo φ(n).
	d, err := inverse(*e, phiN)
	if err != nil {
		fmt.Printf("cannot get modular multiplicative inverse of e modulo φ(n): %v\n", err)
		os.Exit(1)
	}
	// now encrypt plaintext with the public key and verify
	// that when decrypting with the private key, we get the
	// same plaintext back.
	ciphertext := encrypt(*plaintext, *e, *n)
	fmt.Printf("Encrypted %d to %d\n", *plaintext, ciphertext)
	decrypted := decrypt(ciphertext, d, *n)
	fmt.Printf("Decrypted %d to %d\n", ciphertext, decrypted)
	if decrypted != *plaintext {
		fmt.Println("RSA encryption/decryption failed!")
		os.Exit(1)
	}
	fmt.Println("RSA encryption/decryption succeeded!")
}
