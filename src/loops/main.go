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
		i++
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

func main() {
	fmt.Println("Loop.")

	n := 5
	m := Factorial(n)
	fmt.Println(m)

	fmt.Println(SumFirstNIntegers(10))

	// fmt.Println(Factorial(-100))
}