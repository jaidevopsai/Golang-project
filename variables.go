package main

import (
	"fmt"
	"time"
)

// Variable samples for Go

// message is a string variable holding a greeting message
var message string = "Hello, Variables!"

// number is an int variable holding a sample integer
var number int = 42

// pi is a float64 variable holding the value of pi
var pi float64 = 3.14 // Short variable declaration
// isActive is a bool variable indicating active status
var isActive bool = true

// x, y, z are int variables holding multiple values
var x, y, z int = 1, 2, 3

// arr is an array of 3 integers
var arr = [3]int{10, 20, 30}

// slice is a slice of strings
var slice = []string{"apple", "banana", "cherry"}

// m is a map from string to int
var m = map[string]int{"foo": 1, "bar": 2}

// b is a byte variable holding the ASCII value of 'A'
var b byte = 'A'

// r is a rune variable holding a Unicode character
var r rune = '世'

// ptr is a pointer to the number variable
var ptr *int = &number

// ch is a buffered channel of int with capacity 1
var ch chan int = make(chan int, 1)

// t is a time.Time variable holding the current time
var t time.Time = time.Now()

// complexNum is a complex128 variable holding a complex number
var complexNum complex128 = 1 + 2i

// arr2d is a 2x2 array of integers
var arr2d = [2][2]int{{1, 2}, {3, 4}}

// interfaceVar is an empty interface holding a string
var interfaceVar interface{} = "I am interface"

// Person is a struct with Name and Age fields
type Person struct {
	Name string
	Age  int
}

// person is a variable of type Person
var person = Person{"Alice", 30}

// greeting is a string constant holding a welcome message
const greeting = "Welcome to Go!"

// PrintAll prints all variables and constants declared in this file
func PrintAllFromVariables() {
	fmt.Println("message is a string variable holding a greeting message")
	fmt.Println("message:", message)
	fmt.Println("number is an int variable holding a sample integer")
	fmt.Println("number:", number)
	fmt.Println("pi is a float64 variable holding the value of pi")
	fmt.Println("pi:", pi)
	fmt.Println("isActive is a bool variable indicating active status")
	fmt.Println("isActive:", isActive)
	fmt.Println("x, y, z are int variables holding multiple values")
	fmt.Println("x, y, z:", x, y, z)
	fmt.Println("arr is an array of 3 integers")
	fmt.Println("arr:", arr)
	fmt.Println("slice is a slice of strings")
	fmt.Println("slice:", slice)
	fmt.Println("m is a map from string to int")
	fmt.Println("m:", m)
	fmt.Println("greeting is a string constant holding a welcome message")
	fmt.Println("greeting:", greeting)
	fmt.Println("b is a byte variable holding the ASCII value of 'A'")
	fmt.Println("b (byte):", b)
	fmt.Println("r is a rune variable holding a Unicode character")
	fmt.Println("r (rune):", r)
	fmt.Println("ptr is a pointer to the number variable")
	fmt.Println("ptr (*int):", ptr, "points to:", *ptr)
	fmt.Println("ch is a buffered channel of int with capacity 1")
	fmt.Println("ch (chan int):", ch)
	fmt.Println("t is a time.Time variable holding the current time")
	fmt.Println("t (time.Time):", t)
	fmt.Println("complexNum is a complex128 variable holding a complex number")
	fmt.Println("complexNum (complex128):", complexNum)
	fmt.Println("arr2d is a 2x2 array of integers")
	fmt.Println("arr2d ([2][2]int):", arr2d)
	fmt.Println("interfaceVar is an empty interface holding a string")
	fmt.Println("interfaceVar (interface{}):", interfaceVar)
	fmt.Println("person is a variable of type Person")
	fmt.Println("person (struct):", person)
}
