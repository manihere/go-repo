package main

import (
	"fmt"
)

var (
	a = 1 + 1
	b = 2
)

func main() {

	fmt.Printf("1 + 1= %d \n", a)
	fmt.Printf("a = %d \n", a)
	fmt.Printf("b = %d \n", b)
	fmt.Printf("a + 20 = %d \n", a+20)
	fmt.Printf("a + b = %d \n", a+b)
	//, a, "b = %d \n", b, "a + b= %d \n", a+b)

}
