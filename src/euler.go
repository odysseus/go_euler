package main

import (
	"fmt"
	"time"
	"strconv"
	"os"
	"bufio"
)

func main() {
	t := time.Now()

	fmt.Printf("Project Euler Problem 1: %d\n", euler1())
	fmt.Printf("Project Euler Problem 2: %d\n", euler2())
	fmt.Printf("Project Euler Problem 3: %d\n", euler3())
	fmt.Printf("Project Euler Problem 4: %s\n", euler4())
	fmt.Printf("Project Euler Problem 5: %d\n", euler5())
	fmt.Printf("Project Euler Problem 6: %d\n", euler6())
	fmt.Printf("Project Euler Problem 7: %d\n", euler7())
	fmt.Printf("Project Euler Problem 8: %d\n", euler8())
	fmt.Printf("Project Euler Problem 9: %d\n", euler9())
	fmt.Printf("Project Euler Problem 10: %d\n", euler10())
	
	fmt.Printf("Took: %0.3fs\n", time.Since(t).Seconds())
}

////
// HELPER METHODS
////

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func primeSieve(n int) []bool {
	var sieve = make([]bool, (n+1))
	for i:=2; i<len(sieve); i++ {
		sieve[i] = true
	}
	for i:=2; i*i <= n; i++ {
		if sieve[i] {
			j := 2
			for i * j <= n {
				sieve[i*j] = false
				j++
			}
		}
	}
	return sieve
}

// 50 Millionth Prime in ~17s
func primeAt(n int) int {
	size := (n / 50) * 1000
	sieve := primeSieve(size)
	count := 0
	i := 0
	for count < n {
		i++
		if sieve[i] {
			count++
		}
	}
	return i
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isPalindrome(s string) bool {
	return s == reverse(s)
}

func findTrips(m, n int) []int {
	a := m*m - n*n
	b := 2*m*n
	c := m*m + n*n
	return []int{a, b, c}
}

func pyTrips(m, n int) []int {
	if m == n { 
		return []int{0,0,0} 
	} else if m > n {
		return findTrips(m,n)
	} else {
		return findTrips(n,m)
	}
}

func arrSum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

////
// EULER PROBLEMS
////

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

// Find the largest prime factor of 600851475143
func euler3() int {
	n := 600851475143
	sieve := primeSieve(10000)
	for i, v := range sieve {
		if n == i {
			break
		}
		if v && n%i == 0 {
			for n%i == 0 {
				n /= i
			}
		}
	}
	return n
}

// Find the largest palindrome made by the product of two 3-digit numbers
func euler4() string {
	max := 0
	for i:=0; i < 1000; i++ {
		for x:=0; x < 1000; x++ {
			product := x * i
			if product > max && isPalindrome(strconv.Itoa(product)) {
				max = product
			}
		}
	}
	return strconv.Itoa(max)
}

// What is the smallest number that is evenly divisble by
// all the numbers from 1 to 20
func euler5() int {
	test := 20
	finished := false
	for !finished {
		finished = true
		test += 20
		for i:=11; i<=20; i++ {
			if test % i != 0 {
				finished = false
				break
			}
		}
	}
	return test
}

// Find the difference between the sum of squares and 
// the square of sums for the first 100 natural numbers
func euler6() int {
	sumsquares := 0
	squaresums := 0
	for i:=1; i<101; i++ {
		sumsquares += i*i
		squaresums += i
	}
	squaresums *= squaresums
	return squaresums - sumsquares
}

// Find the 10,001st prime
func euler7() int {
	return primeAt(10001)
}

// Find the largest sum made by 13 consecutive digits in
// a 1,000 digit number
func euler8() int64 {
	// Read the file
	f, err := readLines("./euler8.txt")
	if err != nil { panic(err) }
	// Concatenate the lines into a single string
	bigstring := ""
	for _, l := range f {
		bigstring += l
	}
	// Convert each character to an int
	var ints [1000]int
	for i:=0; i<len(bigstring); i++ {
		ints[i], err = strconv.Atoi(bigstring[i:i+1])
	}
	// Finally, calculate the max
	var max int64 = 0
	for i:=0; i<(len(ints)-12); i++ {
		var total int64 = 1
		for x:=0; x<13; x++ {
			total *= int64(ints[i+x])
		}
		if total > max { max = total }
	}
	return max
}

// Find the only Pythagorean triplet whose sum is equal
// to 1,000 and return the product of a, b, and c
func euler9() int {
	for i:=0; i<100; i++ {
		for y:=0; y<100; y++ {
			var trips []int = pyTrips(i,y)
			if arrSum(trips) == 1000 {
				return trips[0] * trips[1] * trips[2]
			}
		}
	}
	return 0
}

// Find the sum of all primes below two million
func euler10() int64 {
	sieve := primeSieve(2000000)
	var total int64 = 0
	for i, v := range sieve {
		if v { total += int64(i) }
	}
	return total
}










