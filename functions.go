package main

import (
	"fmt"
)

func hello() { //simple function
	fmt.Println("Function hello() was executed")
}

func param(name string, age int) { //func with parameter
	fmt.Println("Hello", name, ", you are aged ", age)
}

func returnoutput(a int, b int) int { //we define the return value type before {
	return a + b

}

func multiReturn(a int, y string) (total int, txt string) { //return multiple values
	total = a + 5
	txt = "hello " + y
	return
}

func main() {
	hello()
	hello()

	param("Manish", 30)
	param("Panwar", 15)

	fmt.Println("a+b=", returnoutput(2, 3))

	total := returnoutput(5, 5) //storing return value in var
	fmt.Println("Total of 5+ 5 =", total)

	fmt.Println(multiReturn(5, "Manish"))

	a, b := multiReturn(10, "Manish") //storing return values in var
	fmt.Println(a, b)

}
