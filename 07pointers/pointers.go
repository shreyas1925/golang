package main

import "fmt"

func main() {

	// Declare a pointer variable

	// var myptr *int

	// myptr = 10

	// fmt.Println("Address of ptr: ", &myptr)

	number := 10

	var myptr = &number

	fmt.Println("Value is: ", *myptr)
	fmt.Println("Address is: ", &number)

	*myptr = *myptr * 20
	fmt.Println("Value is: ", *myptr)
	fmt.Println("Value is: ", number)

}
