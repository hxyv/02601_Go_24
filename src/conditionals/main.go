package main

import "fmt"

// Min takes two integers as input and reutrns their minimum.
func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//  WhichIsGreater takes two integers as input. It reutrns 0 if two input integers are equal, 1 if the first input integer is larger, and -1 if the second input integer is larger.
func WhichIsGreater(x, y int) int {
	if x == y { // == means testing whether two expressions are equal
		return 0
	} else if x > y {
		return 1
	} else { // we can infer x < y
		return -1
	}
}

// SameSign takes two integers as input. It returns true if the numbers have the same sign and false otherwise, assuming that 0 has the same sign as all integers.
func SameSign(x, y int) bool {
	if x * y >= 0 { // >= is "greater than or equal"
		return true
	}
	// If we make it here, then we know that x * y < 0
	return false

	// Optimize by "return x * y >= 0"
}

// PositiveDifference takes two integers as input and returns the absolute value of the difference between the two integers.
func PositiveDifference(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return a - b
	} else { // we can infer that a < b
		return b - a
	}
}

// Implement PositiveDifference using only two cases
func PositiveDifference2(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

// PositiveDifference3 introduce a varaible c to store the value that eventually return at the end of the function in the case that a and b are not equal.
func PositiveDifference3(a, b int) int {
	var c int // has scope of the entire function
	if a == b {
		return 0
	} else if a > b {
		c = a - b
	} else { // we can infer that a < b
		c = b - a}

	return c
}

func main() {
	fmt.Println("The minimum of 3 and 4 is", Min2(3, 4))

	fmt.Println(WhichIsGreater(3, 5))
	fmt.Println(WhichIsGreater(42, 42))
	fmt.Println(WhichIsGreater(-2, -7))

	fmt.Println(SameSign(1, 2))
	fmt.Println(SameSign(1, -1))

	fmt.Println(PositiveDifference(5, 2))
	fmt.Println(PositiveDifference2(5, 2))
	fmt.Println(PositiveDifference(2, 5))
	fmt.Println(PositiveDifference2(2, 5))
	fmt.Println(PositiveDifference(2, 2))
	fmt.Println(PositiveDifference2(2, 2))

	fmt.Println(PositiveDifference3(5, 2))
	fmt.Println(PositiveDifference3(2, 5))
	fmt.Println(PositiveDifference3(2, 2))
}