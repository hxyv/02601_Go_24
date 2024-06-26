package main

import "fmt"

// Factorial takes an integer n as input and returns n! = n*(n-1)*(n-2)*...*2*1
func Factorial(n int) int {
	if n < 0 { // let's handle negative input
		panic("Error: Negative input given to Factorial()!")
	}
	p := 1
	i := 1
	
	// Go doesn't have a while keyword, and it uses "for" instead
	for i <= n {
		p *= i
		// i++
	}

	return p
}

// SumFirstNIntegers takes an integer n as input and returns the sum of the first n positive integers
func SumFirstNIntegers(n int) int{
	if n < 0 {
		panic("Error: Negative integer given to SumFirstNIntegers()!")
	}

	sum := 0
	i := 1

	for i <= n {
		sum += i
		i++
	}

	return sum
}

// AnotherFactorial takes an integer n as input and returns n! = n*(n-1)(n-2)*...*2*1, using a for loop
func AnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: Negative input given to AnotherFactorial")
	}

	p := 1

	// for every integer i between 1 and n, p = p*i
	for i := n; i >= 1; i-- {
		p *= i
	}

	return p
}

// SumEven takes an integer k as input and returns the sum of every even integer up to and including n
func SumEven(k int) int {
	sum := 0

	// for every even integer j between 1 and k, sum = sum + j
	for j := 2; j <= k; j += 2 {
		sum += j
	}

	return sum
}

// AnotherSum takes an integer n as input and returns the sum of the first n positive integers, using a for loop
func AnotherSum(n int) int {
	if n < 0 {
		panic("Error: negative integer given to AnotherSum")
	}

	sum := 0

	for k := 1; k <= n; k++ {
		sum += k
	}

	return sum
}

func main() {
	fmt.Println("Loop.")

	n := 5
	m := Factorial(n)
	fmt.Println(m)

	fmt.Println(SumFirstNIntegers(10))

	// fmt.Println(Factorial(-100))

	fmt.Println(AnotherFactorial(5))
	fmt.Println(AnotherSum(10))
}