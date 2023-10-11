package main

import (
	"fmt"
)

func main() {
	var x = 4

	switch x { //will run only the condition which is true, BREAK not needed
	//we need to ensure that all conditions have same data type else we will get exception
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid Day!")
	}

	//multicase switch
	var y = 0

	switch y {
	case 1, 3:
		fmt.Println("Odd number")
	case 2, 4:
		fmt.Println("Even number")
	case 0:
		fmt.Println("Number is neither Odd not Even i.e. 0")
	default:
		fmt.Println("Not a number in range 0-4")
	}

}
