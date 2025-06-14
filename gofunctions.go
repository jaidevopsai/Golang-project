package main

import "fmt"

// add returns the sum of two integers
func add(a int, b int) int {
	return a + b
}

// greet prints a greeting message
func greet(name string) {
	fmt.Println("Hello,", name+"!")
}

// isEven returns true if the number is even
func isEven(n int) bool {
	return n%2 == 0
}

// RunFunctionsFromFunctions runs all sample functions for learning
func RunFunctionsFromFunctions() {
	fmt.Println("RunFunctionsFromFunctions: This is a sample function from gofunctions.go!")
	fmt.Println("add(3, 5):", add(3, 5))
	greet("Gopher")
	fmt.Println("isEven(10):", isEven(10))
	fmt.Println("isEven(7):", isEven(7))
}
