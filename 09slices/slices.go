package main

import "fmt"
import "sort"

func main() {
	// fmt.Println("OK Slices")

	guestsNames := []string{}
	fmt.Printf("Type of guestsNames is: %T\n", guestsNames)

	guestsNames = append(guestsNames,"John","Kalyan","David")
	fmt.Println("Following guests are:",guestsNames)
	
	sort.Strings(guestsNames)
	fmt.Println("Following guests are:",guestsNames)

	// while using slices if we declare the size while creating the slices using make keyword we cannot use extra memory through arr[i] 
	// but we can use more elements using append method
	
	luckynumbers := make([]int,4)

	luckynumbers[0] = 10
	luckynumbers[1] = 20
	luckynumbers[2] = 17
	luckynumbers[3] = 40

	// luckynumbers[4] = 75 cannot because we have an length of 4
	fmt.Println(luckynumbers)

	luckynumbers = append(luckynumbers,75,78,19,25,18,16)

	fmt.Println(luckynumbers)

	sort.Ints(luckynumbers)
	fmt.Println(luckynumbers)

	// Remove some particular elements from slices based on indexes and

	learnings := make([]string,4)
	learnings[0] = "ReactJS"
	learnings[1]="Express&Node"
	learnings[2]="FLASK"
	learnings[3]="GITOPS"
	
    var	index int= 2
	learnings = append(learnings[:index],learnings[index+1:]...)
	fmt.Println(learnings)

}