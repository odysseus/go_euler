package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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
	fmt.Printf("Project Euler Problem 11: %d\n", euler11())
	fmt.Printf("Project Euler Problem 12: %d\n", euler12())
	fmt.Printf("Project Euler Problem 13: %s\n", euler13())
	fmt.Printf("Project Euler Problem 14: %d\n", euler14())
	fmt.Printf("Project Euler Problem 15: %d\n", euler15())
	fmt.Printf("Project Euler Problem 16: %d\n", euler16())
	fmt.Printf("Project Euler Problem 17: %d\n", euler17())
	fmt.Printf("Project Euler Problem 18: %d\n", euler18())
	fmt.Printf("Project Euler Problem 19: %d\n", euler19())
	fmt.Printf("Project Euler Problem 20: %d\n", euler20())
	fmt.Printf("Project Euler Problem 21: %d\n", euler21())
	fmt.Printf("Project Euler Problem 22: %d\n", euler22())
	fmt.Printf("Project Euler Problem 23: %d\n", euler23())
	fmt.Printf("Project Euler Problem 24: %s\n", euler24())
	

	fmt.Printf("Took: %0.3fs\n", time.Since(t).Seconds())
}

////
// HELPER METHODS
////

// Global prime sieve avoids initializing this constantly
var Sieve []bool = primeSieve(10000000)

// PRIMES AND FACTORING

// Generates a prime sieve up to n
func primeSieve(n int) []bool {
	var sieve = make([]bool, (n + 1))
	for i := 2; i < len(sieve); i++ {
		sieve[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			j := 2
			for i*j <= n {
				sieve[i*j] = false
				j++
			}
		}
	}
	return sieve
}

// Finds the Nth prime number
// 50 Millionth Prime in ~17s
func primeAt(n int) int {
	sieve := make([]bool, 0, 0)
	if n < len(Sieve) {
		sieve = Sieve
	} else {
		size := (n / 50) * 1000
		sieve = primeSieve(size)
	}
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

// Returns a slice of all prime and non-prime factors for a number
func factors(n int) []int {
	facts := make([]int, 0, 2)
	facts = []int{ 1 }
	r := int(math.Sqrt(float64(n)))
	for i:=2; i<=r; i++ {
		if n % i == 0 {
			facts = append(facts, i)
			facts = append(facts, n/i)
		}
	}
	sort.Ints(facts)
	return facts
}

// Finds only the prime factors for a number
func primeFacts(n int64) []int {
	facts := make([]int, 0, 2)
	r := int(math.Sqrt(float64(n)))
	for i:=0; i<=r; i++ {
		if Sieve[i] && n%int64(i) == 0 {
			facts = append(facts, i)
		}
	}
	return facts
}

// Finds the number of factors for a number, prime and non-prime
func factorsCount(n int) int {
	count := 1
	for i, v := range Sieve {
		if v && n%i == 0 {
			current := 1
			for n%i == 0 {
				n /= i
				current++
			}
			count *= current
		}
		if n == 1 {
			break
		}
	}
	return count-1
}

// Finds the sum of all divisors for n
func sumdiv(n int) int {
	r := int(math.Sqrt(float64(n)))
	total := 1
	for i := 2; i <= r; i++ {
		if n%i == 0 {
			total += i
			total += n / i
		}
	}
	if r*r == n {
		total -= r
	}
	return total
}

func isPerfect(n int) bool {
	return sumdiv(n) == n
}

func isAbundant(n int) bool {
	return sumdiv(n) > n
}

func isDeficient(n int) bool {
	return sumdiv(n) < n
}

// Simple iterative function to find the factorial of n
func factorial(n int) int64 {
	var fact int64 = 1
	for i := n; i > 1; i-- {
		fact *= int64(i)
	}
	return fact
}

// FILES AND UTILITY

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

// Reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Returns true is the string is a palindrome
func isPalindrome(s string) bool {
	return s == reverse(s)
}

// ARRAY OPERATIONS

// Total all values in a slice of ints
func arrSum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

// Tests for slice inclusion
func arrIncludes(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

// PROBLEM SPECIFIC HELPER FUNCTIONS

// Finds a set of pythagorean triplets given two positive integers
// m and n with m > n
func findTrips(m, n int) []int {
	a := m*m - n*n
	b := 2 * m * n
	c := m*m + n*n
	return []int{a, b, c}
}

// Wrapper method that checks for m being larger than n and m != n
// which makes iteratively calling findTrips in a loop possible
func pyTrips(m, n int) []int {
	if m == n {
		return []int{0, 0, 0}
	} else if m > n {
		return findTrips(m, n)
	} else {
		return findTrips(n, m)
	}
}

// Finds the number of characters used to write out a number
// Suitable for numbers below 10,000
func writtenCharCount(i int) int {
	// Tallys the character counts by place value
	var ones []int = []int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4}
	var tens []int = []int{0, 0, 6, 6, 5, 5, 5, 7, 6, 6}
	var hundreds []int = []int{0, 10, 10, 12, 11, 11, 10, 12, 12, 11}
	var thousands []int = []int{0, 11, 11, 13, 12, 12, 11, 13, 12, 12}
	// Teens are a special case and are treated differently
	var teens []int = []int{3, 6, 6, 8, 8, 7, 7, 9, 8, 8}
	total := 0
	s := strconv.Itoa(i)
	// If number is 1000 or greater (Eg. 1234)
	if i/1000 > 0 {
		// Get the number of characters used spelling out the "___ thousand..."
		// part by referencing the array (Eg. [1]234 => thousands[1] = 11)
		total += thousands[int(s[0]-'0')]
		// Set i equal to the remainder below 1000 (Eg. i = 234)
		i %= 1000
		// Set the string to the new value (Eg. s = "234")
		// Do not use substrings, this causes problems when the number has leading zeroes
		// so reconvert it each time
		s = strconv.Itoa(i)
	}
	if i/100 > 0 {
		// Repeat the process above for the hundreds
		total += hundreds[int(s[0]-'0')]
		// add the "...and..." if it is not evenly divisble by 100
		if i%100 != 0 {
			total += 3
		}
		i %= 100
		s = strconv.Itoa(i)
	}
	if i/10 > 0 {
		// For the tens place we need to check for teens first
		if i > 9 && i < 20 {
			// Teens take care of both the tens and the ones spot
			// so we divide by i since no characters remain to be counted
			total += teens[i-10]
			i %= i
		} else {
			// Otherwise the process is the same as above
			total += tens[int(s[0]-'0')]
			i %= 10
			s = strconv.Itoa(i)
		}
	}
	if i/1 > 0 {
		total += ones[int(s[0]-'0')]
	}
	return total
}

// Finds the Nth lexicographic permutation intelligently by inferring
// each subsequent number using the amount of permutations remaining
// combined with the amount of permutations possible for x numbers
// (Possible permutations are found using factorials)
// For example: In the numbers 0..9, there are 362,880 permutations
// for each starting number, so to find the millionth permutation the
// permutation must start with 2, because 362,880 * 2 is less, and
// 362,880 * 3 is more than a million. Following that same logic
// we can compute the value for each subsequent digit
func nthPermutation(series []int, n int) string {
	l := len(series)
	sort.Ints(series)
	var remaining int64 = int64(n) - 1
	permNum := ""
	for i := 1; i < l; i++ {
		j := remaining / factorial(l-i)
		remaining = remaining % factorial(l-i)
		permNum += strconv.Itoa(series[j])
		series = append(series[:j], series[j+1:]...)
		if remaining == 0 {
			break
		}
	}
	// If any numbers remain, append them now
	for _, v := range series {
		permNum += strconv.Itoa(v)
	}
	return permNum
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
func euler3() int64 {
	var n int64 = 600851475143
	var i int64 = 0
	for i<int64(len(Sieve)) {
		if n == i { break }
		if Sieve[i] && n%i == 0 {
			for n%i == 0 {
				n /= i
			}
		}
		i++
	}
	return n
}

// Find the largest palindrome made by the product of two 3-digit numbers
func euler4() string {
	max := 0
	for i := 0; i < 1000; i++ {
		for x := 0; x < 1000; x++ {
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
		for i := 11; i <= 20; i++ {
			if test%i != 0 {
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
	for i := 1; i < 101; i++ {
		sumsquares += i * i
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
	f, _ := readLines("./euler8.txt")
	// Concatenate the lines into a single string
	bigstring := ""
	for _, l := range f {
		bigstring += l
	}
	// Convert each character to an int
	var ints [1000]int
	for i, v := range bigstring {
		// This trick works for getting the int value represented by
		// and ASCII character, rather than the ASCII int value assigned
		// to represent the character
		ints[i] = int(v - '0')
	}
	// Finally, calculate the max
	var max int64 = 0
	for i := 0; i < (len(ints) - 12); i++ {
		var total int64 = 1
		for x := 0; x < 13; x++ {
			total *= int64(ints[i+x])
		}
		if total > max {
			max = total
		}
	}
	return max
}

// Find the only Pythagorean triplet whose sum is equal
// to 1,000 and return the product of a, b, and c
func euler9() int {
	for i := 0; i < 100; i++ {
		for y := 0; y < 100; y++ {
			var trips []int = pyTrips(i, y)
			if arrSum(trips) == 1000 {
				return trips[0] * trips[1] * trips[2]
			}
		}
	}
	return 0
}

// Find the sum of all primes below two million
func euler10() int64 {
	sieve := Sieve[0:2000000]
	var total int64 = 0
	for i, v := range sieve {
		if v {
			total += int64(i)
		}
	}
	return total
}

// Find the largest product of 4 adjacent numbers in a 20x20 grid
func euler11() int {
	// Read the file
	f, _ := readLines("./euler11.txt")
	// Create a multidimensional array to store the ints
	ints := make([][]int, 20)
	// Parse the strings
	for i, v := range f {
		// Fields splits on whitespace and returns a
		// slice containing the individual strings
		rowStrings := strings.Fields(v)
		row := make([]int, 20)
		// Convert all the string values to ints
		for x, y := range rowStrings {
			row[x], _ = strconv.Atoi(y)
		}
		// And add them to the final array
		ints[i] = row
	}
	// Problem Logic
	max := 0
	for i := 0; i < len(ints); i++ {
		for x := 0; x < len(ints[i]); x++ {
			// Left-Right
			if x < 17 {
				product := ints[i][x] * ints[i][x+1] * ints[i][x+2] * ints[i][x+3]
				if product > max {
					max = product
				}
			}
			// Up-Down
			if i < 17 {
				product := ints[i][x] * ints[i+1][x] * ints[i+2][x] * ints[i+3][x]
				if product > max {
					max = product
				}
			}
			// Diagonal Ascending /
			if i > 3 && x < 17 {
				product := ints[i][x] * ints[i-1][x+1] * ints[i-2][x+2] * ints[i-3][x+3]
				if product > max {
					max = product
				}
			}
			// Diagonal Descending \
			if i < 17 && x < 17 {
				product := ints[i][x] * ints[i+1][x+1] * ints[i+2][x+2] * ints[i+3][x+3]
				if product > max {
					max = product
				}
			}
		}
	}
	return max
}

// Find the first triangle number with over 500 factors
func euler12() int {
	n := 0
	triangle := 0
	for true {
		n++
		triangle += n
		if factorsCount(triangle) > 500 {
			break
		}
	}
	return triangle
}

// Find the first ten digits of the sum of 100 50-digit numbers
func euler13() string {
	var result int64 = 0
	f, _ := readLines("./euler13.txt")
	for _, v := range f {
		lineInt, _ := strconv.Atoi(v[0:12])
		result += int64(lineInt)
	}
	return strconv.Itoa(int(result))[0:10]
}

// Find the number, below 1,000,000, with the longest Collatz sequence
func euler14() int {
	longest := 0
	maxn := 0
	for i := 1; i < 1000000; i += 2 {
		var n int64 = int64(i)
		count := 1
		for n > 1 {
			if n%2 == 0 {
				n /= 2
				count++
			} else {
				n = n*3 + 1
				count++
			}
		}
		if count > longest {
			longest = count
			maxn = i
		}
	}
	return maxn
}

// How many possible paths are there through a 20x20 grid
func euler15() *big.Int {
	// The answer is the binomial coefficient for 40 and 20
	return new(big.Int).Binomial(40, 20)
}

// Find the sum of the digits for the number 2**1000
func euler16() int {
	n := big.NewInt(2)
	y := big.NewInt(1000)
	x := new(big.Int).Exp(n, y, nil)
	s := x.String()
	total := 0
	for _, v := range s {
		total += int(v - '0')
	}
	return total
}

// Find the number of letters used spelling out all the
// numbers from one to one thousand
func euler17() int {
	total := 0
	for i := 1; i <= 1000; i++ {
		total += writtenCharCount(i)
	}
	// 21124
	return total
}

// Find the path through the triangle that leads
// to the highest sum
func euler18() int {
	// Read the file
	f, _ := readLines("./euler18.txt")
	// Feed that into a string array
	dataStrings := make([][]string, len(f))
	for i, v := range f {
		dataStrings[i] = strings.Fields(v)
	}
	// Create and parse a nested int array from the strings
	data := make([][]int, len(dataStrings))
	for i, v := range dataStrings {
		data[i] = make([]int, len(v))
		for x, y := range v {
			converted, _ := strconv.Atoi(y)
			data[i][x] = converted
		}
	}
	// Solve the problem from the bottom up where the choices are obvious, then
	// feed the sum into the spot above, by the time you reach the top of the
	// pyramid the value remaining will be the maximum
	for x := len(data) - 2; x >= 0; x-- {
		for i, _ := range data[x] {
			if data[x+1][i] > data[x+1][i+1] {
				data[x][i] += data[x+1][i]
			} else {
				data[x][i] += data[x+1][i+1]
			}
		}
	}
	return data[0][0]
}

// Find the number of months during the 20th century that began with a Sunday
func euler19() int {
	total := 0
	for y := 1901; y < 2001; y++ {
		for m := 1; m < 13; m++ {
			d := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
			if d.Weekday() == time.Sunday {
				total += 1
			}
		}
	}
	return total
}

// Find the sum of the digits for the number 100!
func euler20() int {
	n := big.NewInt(100)
	for i := 99; i > 1; i-- {
		x := big.NewInt(int64(i))
		n = new(big.Int).Mul(n, x)
	}
	s := n.String()
	total := 0
	for _, v := range s {
		total += int(v - '0')
	}
	return total
}

// Find the sum of all amicable numbers below 1000
func euler21() int {
	amic := make([]int, 0, 10)
	for i := 2; i < 10000; i++ {
		if i == sumdiv(sumdiv(i)) && sumdiv(i) != i {
			if !arrIncludes(amic, i) {
				amic = append(amic, i, sumdiv(i))
			}
		}
	}
	total := 0
	for _, v := range amic {
		total += v
	}
	return total
}

// Find the name scores of all names in the .txt file
func euler22() int64 {
	// Read the file
	f, _ := readLines("names.txt")
	// Split into an array
	names := strings.Split(f[0], ",")
	// Remove the quotes from the individual strings
	for i, v := range names {
		names[i] = strings.Trim(v, "\"")
	}

	// Sort the array
	sort.Strings(names)
	var total int64 = 0
	// Then iterate the array and sum up the scores
	for i, v := range names {
		// Convert the string to a byte array
		byts := []byte(v)
		score := 0
		for _, val := range byts {
			// Rather than using an array to find the letter scores, we can
			// use the ASCII bytecode. Capital A is 65, B is 66, etc. so we
			// just need to subtract the byte value by 64
			score += int(val) - 64
		}
		// Compute the final score by multiplying the index + 1 for the rank
		// and the sum of the byte array for the letter score
		total += int64(score) * int64(i+1)
	}
	// And we're done
	return total
}

// Find the sum of all numbers that cannot be written as the sum of
// two abundant numbers
func euler23() int {
	abun := make([]int, 0, 10)
	upperRange := 28124
	for i := 2; i < upperRange; i++ {
		if isAbundant(i) {
			abun = append(abun, i)
		}
	}
	sieve := make([]bool, upperRange, upperRange)
	for i, _ := range sieve {
		sieve[i] = true
	}
	for _, v1 := range abun {
		for _, v2 := range abun {
			if v1+v2 <= upperRange-1 {
				sieve[v1+v2] = false
			}
		}
	}
	total := 0
	for i, v := range sieve {
		if v {
			total += i
		}
	}
	return total
}

func euler24() string {
	series := make([]int, 10, 10)
	series = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return nthPermutation(series, 1000000)
}
