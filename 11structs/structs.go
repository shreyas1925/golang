package main

import "fmt"

func main() {
	// fmt.Println("Structs in golang")

	shreyas := Student{"Shreyas", "BE", "4NI19IS094", "Male"}
	fmt.Printf("Students details are as follows : %+v\n", shreyas)
	fmt.Printf("%v has a USN of %v", shreyas.Name, shreyas.USN)

}

type Student struct {
	Name   string
	Class  string
	USN    string
	Gender string
}
