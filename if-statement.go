package main

import (
	"fmt"
)

var (
	y    = 20
	x    = 10
	time = 12
)

func main() {

	//IF-ELSE, will execute statement block for which  condition is true
	if x > y {
		fmt.Println("x is greater than y") //run if true
	} else { //else should be just after } as shown, else it will throw an error
		fmt.Println("x is less than y")
	}

	//Else-If, will execute the block whose condition is false
	if x > y {
		fmt.Println(" x > y") //runs if condition1 is True
	} else if x < y {
		fmt.Println(" x < y ") //runs if condition1 is False, but condition2 is true
	} else {
		fmt.Println("x and y are equal") //runs if both conditions are false
	}

	//if two conditions are true, then only the first block is run
	if x <= 20 {
		fmt.Println("x <= 20") //true, will print
	} else if x < 30 {
		fmt.Println("x is less than 30") //true, will not print
	} else {
		fmt.Println("x < 40") //will not print since condition1 or condition2 were false
	}

	//Nested If- loop inside loop
	if x < 15 {
		fmt.Println("x < 15")
		if x < 25 {
			fmt.Println("x is also less than 25")
		}
	} else {
		fmt.Println(x < 10)
	}
}
