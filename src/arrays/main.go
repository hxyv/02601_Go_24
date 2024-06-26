package main

import ("fmt")

func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: negative input given to FactorialArray().")
	}

	fact := make([]int, n+1)
	fact[0] = 1

	for k := 1; k <= n; k++ {
		fact[k] = fact[k-1]*k
	}

	return fact
}

func ChangeFirstElementArray(a [5]int) {
    a[0] = 1
}
func ChangeFirstElementSlice(a []int) {
    a[0] = 1
}

// MinIntegerArray takes as input a slice of integers.
// It returns the minimum value in the array.
func MinIntegerArray(list []int) int {
    if len(list) == 0 {
        panic("Error: Empty list given as input.")
    }

    var m int // default value = 0
	
    // range over list, updating m if we find bigger value
    for i, val := range list {
        if i == 0 || val < m {
            m = val
        }
    }
    return m
}

// MinIntegers takes as input an arbitrary collection of integers.
// It returns the minimum value among the input integers.
func MinIntegers(numbers ...int) int {
    if len(numbers) == 0 {
        panic("Error: No integers given as input.")
    }

    return MinIntegerArray(numbers)
}

func main() {
	fmt.Println("Arrays (and slices).")

	// Arrays in Go have a fixed, constant size
	var list [6]int // gives 6 default (zero) values
	list[0] = -8
	i := 3
	list[2*i-4] = 17 // 2*i - 4 = 2
	list[len(list)-1] = 43
	fmt.Println(list)

	var a []int
	n := 4
	a = make([]int, n)
	a[0] = 42
	a[2] = -17
	fmt.Println(a)

	b := make([]int, n+2)
	fmt.Println(b)

	var c [5]int // is equal to [0, 0, 0, 0, 0]
    d := make([]int, 5) // also is equal to [0, 0, 0, 0, 0]
    ChangeFirstElementArray(c)
    ChangeFirstElementSlice(d)
    fmt.Println(c) // prints [0, 0, 0, 0, 0]
    fmt.Println(d) // prints [1, 0, 0, 0, 0]
}
