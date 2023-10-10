package main

import (
	"fmt"
)

var (
	Array1 = []int{1, 2, 3, 4, 5, 6}
	slice0 = []int{}        //empty slice
	slice1 = []int{1, 2, 3} //slice like an array
	slice2 = Array1[2:4]    //slice made from array

	//using make() to create slices
	slice3 = make([]int, 5, 10) //make([]type, len, cap)
	slice4 = make([]int, 5)     //here cap is not defined so cap=len
)

func main() {
	//SLICE is similar to arrays but are powerful
	//slice can also be used to slice a part of array
	fmt.Println(slice0)
	fmt.Println(slice1)

	fmt.Println("slice2 is ", slice2)                   //returns two value in pos 2-4 of array1
	fmt.Printf("Slice2 length is %d ,", len(slice2))    // returns 2 since there are two values in slice
	fmt.Printf("Slice2 capacity is %d /n", cap(slice2)) //returns 4 since array was cap 6, slice starts from pos 2 and goes till end of array 6-2=4

	fmt.Println(slice3)
	fmt.Printf("with cap defined slice3 Length: %d Capacity %d \n", len(slice3), cap(slice3))

	fmt.Printf("Without cap defined cap=len for slice4 Length %d Capacity %d \n", len(slice4), cap(slice4))

	fmt.Println(slice1[1]) //printing particular value

	slice1[2] = 10 //overriding slice value
	fmt.Println(slice1)

	slice1 = append(slice1, 20, 30) //appending values to slice
	fmt.Println(slice1)

	slice1 = append(slice1, slice2...) //... is necessary when appending slices
}
