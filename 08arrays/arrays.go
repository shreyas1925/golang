package main

import "fmt"

func main() {
	fmt.Println("Arrays in golang!!")

	var arr [2]int

	arr[0] = 9
	arr[1] = 5

	fmt.Println("Arrays ", arr)

	luckynumbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Luckynumbers are: ", luckynumbers)

	magic := [3]string{"one", "three"}
	fmt.Println("Magic includes: ", magic)

	// Weird thing in golang 😒

	fmt.Println("Length of magic: ", len(magic)) //Inspite of only 2 elements it still gives the length of actual magic

}
