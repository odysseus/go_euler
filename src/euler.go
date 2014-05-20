package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Project Euler Problem 1: %d\n", euler1())
	fmt.Printf("Project Euler Problem 2: %d\n", euler2())
}

// Find the sum of all multiples of 3 or 5 below 1000
func euler1() int {
	r := 0
	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			r += i
		}
	}
	return r
}

// Find the sum of even valued Fibonacci terms below 4 million
func euler2() int {
	a := 0
	b := 1
	total := 0
	for b < 4000000 {
		if b%2 == 0 {
			total += b
		}
		b = a + b
		a = b - a
	}
	return total
}

