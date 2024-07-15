package main

import (
	"fmt"
	"math"
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

func FibonacciArray(n int) []int {
	fibonacciList := make([]int, 0)

	if n == 0 {
		fibonacciList = append(fibonacciList, 1)
		return fibonacciList
	}

	for i := 0; i <= n; i++ {
		if i == 0 {
			fibonacciList = append(fibonacciList, 1)
		} else if i == 1 {
			fibonacciList = append(fibonacciList, 1)
		} else {
			sum := fibonacciList[i-2] + fibonacciList[i-1]
			fibonacciList = append(fibonacciList, sum)
		}
	}

	return fibonacciList
}

func DividesAll(a []int, d int) bool {
	if d == 0 {
		return false
	}

	for _, value := range a {
		if value % d != 0 {
			return false
		}
	}

	return true
}

func MaxIntegerArray(list []int) int {
	max := list[0]

	for _, value := range list {
		if value > max {
			max = value
		}
	}

	return max
}

func MaxIntegers(numbers ...int) int {
	return MaxIntegerArray(numbers)
}

func MinIntegerArray(list []int) int {
	min := list[0]

	for _, value := range list {
		if value < min {
			min = value
		}
	}

	return min
}

func MinIntegers(numbers ...int) int {
	return MinIntegerArray(numbers)
}

func SumIntegers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}

func GCDArray(a []int) int {
	divisor := 1
	min := MinIntegerArray(a)

	for p := 1; p <= min; p++ {
		if DividesAll(a, p) == true {
			divisor = p
		}
	}

	return divisor
}

func IsPerfect(n int) bool {
	if n == SumProperDivisors(n) {
		return true
	}

	return false
}

func NextPerfectNumber(n int) int {
	if n == 0 {
		return 6
	}

	i := n + 1

	for IsPerfect(i) == false {
		i++
	}

	return i
}

func IsPrime(p int) bool {
    for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
        if p%k == 0 {
            return false
        }
    }
    return true
}

func ListMersennePrimes(n int) []int {
	MersennerList := make([]int, 0)

	for i := 2; i <= n; i++ {
		num := Power(2, i) - 1
		if IsPrime(num) {
			MersennerList = append(MersennerList, num)
		}
	}
	
	return MersennerList
}

func NextTwinPrimes(n int) (int, int) {
	var i int

	if n == 0 {
		i = 3
	} else {
		i = n + 1
	}

	for !IsPrime(i) || !IsPrime(i + 2) {
		i++
	}

	return i, i + 2
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
	fmt.Println(FibonacciArray(0))
	fmt.Println(FibonacciArray(1))
	fmt.Println(FibonacciArray(2))
	fmt.Println(FibonacciArray(3))
	fmt.Println(FibonacciArray(30))

	fmt.Println(DividesAll([]int{2, 4, 6, 8, 10}, 2))
	fmt.Println(DividesAll([]int{1, 3, 5, 7}, 14))
	fmt.Println(DividesAll([]int{104, 8, 19, 203}, 1))
	fmt.Println(DividesAll([]int{-16, -24, -400, -60}, 4))
	fmt.Println(DividesAll([]int{9, 15, 12}, -3))
	fmt.Println(DividesAll([]int{7, 14, 21}, 0))
	fmt.Println(DividesAll([]int{3, 6, 24, 12}, 2))
	fmt.Println(DividesAll([]int{4, 48, 12, 6}, 12))

	fmt.Println(MaxIntegerArray([]int{-47}))
	fmt.Println(MaxIntegerArray([]int{1,2,3,4,5,6,7}))
	fmt.Println(MaxIntegerArray([]int{7,6,5,4,3,2,1}))
	fmt.Println(MaxIntegerArray([]int{-100,-100,-1,-100,-100}))

	fmt.Println(MaxIntegers(-47))
	fmt.Println(MaxIntegers(1,2,3,4,5,6,7))
	fmt.Println(MaxIntegers(7,6,5,4,3,2,1))
	fmt.Println(MaxIntegers(-100,-100,-1,-100,-100))

	fmt.Println(SumIntegers(42))
	fmt.Println(SumIntegers(0, 1, -1))
	fmt.Println(SumIntegers(1, 2, 3, 4, 5, 6))
	fmt.Println(SumIntegers(-1,-1,-1,-1,-1,-1,-1,-1,-1,-1))

	fmt.Println(GCDArray([]int{7}))
	fmt.Println(GCDArray([]int{2,2,2,2,2}))
	fmt.Println(GCDArray([]int{100,200,1,300,400}))
	fmt.Println(GCDArray([]int{12,15,9,6,24}))
	fmt.Println(GCDArray([]int{6,9,16}))

	// Chapter 2.4
	fmt.Println(IsPerfect(1))
	fmt.Println(IsPerfect(2))
	fmt.Println(IsPerfect(6))
	fmt.Println(IsPerfect(10))
	fmt.Println(IsPerfect(14))
	fmt.Println(IsPerfect(28))
	fmt.Println(IsPerfect(300))
	fmt.Println(IsPerfect(496))
	fmt.Println(IsPerfect(8000))
	fmt.Println(IsPerfect(8128))

	fmt.Println(NextPerfectNumber(0))
	fmt.Println(NextPerfectNumber(5))
	fmt.Println(NextPerfectNumber(6))
	fmt.Println(NextPerfectNumber(27))
	fmt.Println(NextPerfectNumber(28))
	fmt.Println(NextPerfectNumber(495))
	fmt.Println(NextPerfectNumber(496))

	fmt.Println(ListMersennePrimes(2))
	fmt.Println(ListMersennePrimes(6))
	fmt.Println(ListMersennePrimes(7))
	fmt.Println(ListMersennePrimes(60))
	fmt.Println(ListMersennePrimes(61))

	// Chapter 2.5
	fmt.Println(NextTwinPrimes(0))
	fmt.Println(NextTwinPrimes(3))
	fmt.Println(NextTwinPrimes(29))
	fmt.Println(NextTwinPrimes(616))
}