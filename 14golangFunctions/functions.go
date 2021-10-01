package main

import "fmt"

func main() {

	welcome("Shreyas", "Golang")
	// fmt.Println("Let's go through the go functions 😉")

	// result := golangAdder(4, 5)
	// result := golangAdder(4, 5, 19, 25, -18, -16)
	result := golangAdder(4, 5, -5, -7, 8, 99, -15)
	fmt.Println(result)

}

func welcome(name string, course string) {

	fmt.Printf("Welome %v for %v learning path\n", name, course)

}

// we can also return more than one return types in the functions

func golangAdder(numbers ...int) int {

	res := 0

	for _, ele := range numbers {
		res += ele
	}

	return res

}
