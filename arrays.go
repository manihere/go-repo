package main

import (
	"fmt"
)

func main() {

	var (
		Array1 = [3]int{1, 2, 3}         //length defined
		Array2 = [...]int{1, 2, 3, 4, 5} //length not defined, compiler decides the length of array
		Array3 = [3]string{"Hi", "There", "buddy"}
		Array4 = [5]int{1, 2, 3}      //partially initialized array
		Array5 = [5]int{1: 10, 2: 20} //initializing only 1&2 position in the array
	)

	fmt.Println(Array1, "\n", Array2, "\n", Array3, "\n")
	Array1[2] = 10 //modify array value
	fmt.Println("Array 1 is now: ", Array1, "\n")
	fmt.Println(Array4)
	fmt.Println(Array5)
	fmt.Println(len(Array5))
}
