package main

import "fmt"

func main() {

	// fmt.Println("What's the syntax of go loops")

	employees := []string{"Arun", "Bhavesh", "Manoj", "Jatin", "Karan"}
	fmt.Println("Employees details: ", employees)

	// Let's try different ways of looping through a slices

	// for emp := 0; emp < len(employees); emp++ {
	// 	fmt.Println(employees[emp])
	// }

	// for emp := range employees {
	// 	fmt.Println(employees[emp])
	// }

	// Let's see foreach kind of looping in golang

	for _, emp := range employees {
		fmt.Println(emp)
	}

}
