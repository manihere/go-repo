// struct is like array, but can keep values of different data types
package main

import (
	"fmt"
)

type Person struct {
	name         string
	age          int
	salary       int
	organization string
}

func main() {
	var pers1 Person
	var pers2 Person

	pers1.name = "Manish"
	pers1.age = 30
	pers1.salary = 10
	pers1.organization = "ABC"

	fmt.Println(pers1.name, pers1.age, pers1.salary, pers1.organization)

	pers2.name = "Panwar"
	pers2.age = 20
	pers2.salary = 15
	pers2.organization = "ABCD"

	fmt.Println(pers2.name, pers2.age, pers2.salary, pers2.organization)

}
