package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 1, 3, 5, 7
	var x, y = 6, "hello" //we can declare & initialize variable of diff kind in same line
	var (
		g = 10 //8-bit
		h = "World"
	) //using var block for more visibility
	const (
		pi      = "3.14" //constant, value cannot be changed)
		A  int  = -100   // variables are case sensitive
		B  uint = 200    //can keep only unsigned values i.e. only positive values
	)
	//pi = 4 will give error "cannot assign to pi (neither addressable nor a map index expression)"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x, y)
	fmt.Println(g, h)
	fmt.Println(pi, "\n", A, B)
	fmt.Printf("var pi has value %v & pi and type is %T", pi, pi)
	/*
		The Printf() function first formats its argument based on the given formatting verb and then prints them.
		Here we will use two formatting verbs:
		%v is used to print the value of the arguments
		%T is used to print the type of the arguments */

	var (
		T   bool = true //The default value of a boolean data type is false
		num int8 = 1    //8-bit to save value, can store values in range2^8 (-128 to +128)
		//you can use int, int8/16/32/64
		pos uint8   = 1 //8-bit unsigned var to save value, can store values in range2^8 (0 to 256)
		flt float32 = 3.3334
		str string  = "Hi There!"
	)
	fmt.Println("\n", T, num, flt, pos, "\n", str)

	
}
