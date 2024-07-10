package main

import (
	"fmt"
)

func Permutation(n, k int) int {
	if k == 0 {
		return 1
	}

	var permutation = n

	for i := n - 1; i > n - k; i-- {
		permutation *= i
	}

	return permutation
}

func Combination(n, k int) int {
	if k == 0 || n == k {
		return 1
	}

	return Permutation(n, k) / Permutation(k, k)
}

func Power(a, b int) int {
	if b == 0 {
		return 1
	}

	power := a

	for i := 1; i < b; i++ {
		power *= a
	}

	return power
}


func SumProperDivisors(n int) int {
	sum := 0

	for i := 1; i < n; i++ {
		if n % i == 0 {
			sum += i
		}
	}

	return sum
}

func FibonacciArray(n int) int {
	fibonacciList := make([]int, 0)
}

func main() {
	// Chapter 2.1
	fmt.Println(Permutation(12, 1))
	fmt.Println(Permutation(6, 6))
	fmt.Println(Permutation(10, 4))
	fmt.Println(Permutation(10, 0))

	fmt.Println(Combination(10, 0))
	fmt.Println(Combination(10, 10))
	fmt.Println(Combination(12, 1))
	fmt.Println(Combination(8, 7))
	fmt.Println(Combination(10, 4))
	fmt.Println(Combination(10, 6))

	// Chapter 2.2
	fmt.Println(Power(2, 3))
	fmt.Println(Power(1, 100))
	fmt.Println(Power(7856, 0))
	fmt.Println(Power(-1, 3))
	fmt.Println(Power(7000, 1))
	fmt.Println(Power(5, 4))
	fmt.Println(Power(10, 6))

	fmt.Println(SumProperDivisors(6))
	fmt.Println(SumProperDivisors(40))
	fmt.Println(SumProperDivisors(2003))
	fmt.Println(SumProperDivisors(1))

	// Chapter 2.3

}