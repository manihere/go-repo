// struct is like array, but can keep values of different data types
package main

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

	pers1.name = "Panwar"
	pers1.age = 20
	pers1.salary = 15
	pers1.organization = "ABCD"
}
