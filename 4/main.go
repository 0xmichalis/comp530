package main

import "fmt"

func gcd(a, b int) int {
	x, y := a, b

	for y != 0 {
		r := x % y
		x = y
		y = r
	}

	return x
}

func main() {
	fmt.Println(gcd(499017086208, 676126714752))
	fmt.Println(gcd(5988737349, 578354589))
}
