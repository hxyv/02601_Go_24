package main

import "fmt"

// SumToInts takes two integers and returns their sum.
func SumToInts(a, b int) int {
	return a + b
}

// DoubleAndDuplicate takes a float64 x as an input.
// It returns two copies of 2x output.
func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0 * x, 2.0 * x
}

// Pi takes no inputs.
// It returns the value 3.14
func Pi() float64 {
	return 3.14
}

// PrintHi simply prints the message "Hi" to the console.
func PrintHi() {
	fmt.Println("Hi")
}

// AddOne takes an integer k as input and returns the value of k + 1
// Notice of the difference between "pass by value" and "pass by reference"
func AddOne(k int) int {
	k = k + 1
	return k
}

func main() {
	x := 3
	n := SumToInts(x, x)
	fmt.Println(n)

	fmt.Println(Pi())
	PrintHi()

	m := 17
	fmt.Println(AddOne(m))
	fmt.Println(m)

	a := 1
	b := 1
	fmt.Println(SumToInts(a,b))
}