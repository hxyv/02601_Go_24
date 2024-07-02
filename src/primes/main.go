package main

import (
	"fmt"
	"math"
	"log"
	"time"
)

// TrivialPrimeFinder takes an integer n and produces a boolean array of length n+1
// where primeBooleans[p] is true if p is prime (and false otherwise)
func TrivialPrimeFinder(n int) []bool {
	var primeBooleans []bool // type is "slice of booleans", which  default to false
	primeBooleans = make([]bool, n+1)

	// in gernal, you'll just use primeBooleans := make([]bool, n+1)
	for p := 2; p <= n; p++ {
		if IsPrime(p) == true {
			primeBooleans[p] = true
		}
	}

	return primeBooleans
}

func IsPrime(p int) bool {
    for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
        if p%k == 0 {
            return false
        }
    }
    return true
}

// SieveOfEratosthenes takes an integer n and returns a slice of n+1
// booleans primeBooleans where primeBooleans[p] is true if p is prime
// and false otherwise.
// It implements the sieve of Eratosthenes approach for finding primes.
func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)

	// Set everything to prime other than primeBooleans[0] and primeBooleans[1]
	for k := 2; k <= n; k++ {
		primeBooleans[k] = true
	}

	// Now, range over primeBooleans, and cross off multiples of first prime we see, iterating this process.
	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		if primeBooleans[p] == true {
			primeBooleans = CrossOffMultiples(primeBooleans, p)
		}
	}
	return primeBooleans
}

//CrossOffMultiples takes a boolean slice and an integer p.  It crosses off
//multiples of p, meaning turning these multiples to false in the slice.
//It returns the slice obtained as a result.
func CrossOffMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1
    for k := 2*p; k <= n; k += p {
		primeBooleans[k] = false // k is composite
	}

	return primeBooleans
}

func ListPrimes(n int) []int {
	primeList := make([]int, 0)
	primeBooleans := SieveOfEratosthenes(n) // remember: primeBooleans will have length n+1
	
	// Range over primeBooleans, and append p to primeList if p is prime
	for p := range primeBooleans {
		if primeBooleans[p] == true {
			primeList = append(primeList, p)
		}
	}

	return primeList
}


func main() {
	fmt.Println("Finding primes.")

	primeBooleans1 := TrivialPrimeFinder(10)
	fmt.Println(primeBooleans1)

	primeBooleans2 := SieveOfEratosthenes(10)
    fmt.Println(primeBooleans2)

	primeList := ListPrimes(11)
    fmt.Println(primeList)

	n := 1000000

    start := time.Now()
    TrivialPrimeFinder(n)
    elapsed := time.Since(start)
    log.Printf("TrivialPrimeFinder took %s", elapsed)
	
    start2 := time.Now()
    SieveOfEratosthenes(n)
    elapsed2 := time.Since(start2)
    log.Printf("SieveOfEratosthenes took %s", elapsed2)
}