package main

import "fmt"

// Problem 1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func toAdd(num int) int {
	if num%3 == 0 || num%5 == 0 {
		return num
	}
	return 0
}

func main() {
	counter := 0
	for i := 0; i <= 1000; i++ {
		counter += toAdd(i)
	}
	fmt.Println("The Total is:", counter)
}
