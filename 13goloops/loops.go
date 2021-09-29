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

	randomnumber := 0

	for randomnumber < 10 {

		if randomnumber == 2 {
			goto statement
		}

		if randomnumber == 5 {
			randomnumber++
			continue
		}

		if randomnumber > 7 {
			break
		}
		fmt.Println("The random number is %v", randomnumber)
		randomnumber++
	}

statement:
	fmt.Println("We have reached the milestone of randomnumber list ✌️")

}
