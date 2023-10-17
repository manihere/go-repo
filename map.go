/*Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.*/

package main

import (
	"fmt"
)

func main() {

	//var a = map[keyType]valueType{key1:value1, key2:value2.....}

	var person1 = map[string]int{"Age1": 20, "Salary": 30, "Experience": 10}

	fmt.Println("Person1 ", person1)

	var Person2 = map[string]string{"Name": "Manish", "Org": "ABC", "lastName": "Panwar"}

	fmt.Println("Person2", Person2)
	fmt.Println(Person2)

	//delete particular value
	delete(Person2, "lastName")
	fmt.Println(Person2)

	//create map using make()
	var c = make(map[string]int)
	c["India"] = +91
	c["Pin"] = 110096
	c["Population in billion"] = 1

	fmt.Println("Facts about India \n %v", c)

	var em = make(map[string]string) //empty map
	fmt.Println(em)

	fmt.Println(c["India"]) // print particular value
	c["Pin"] = 110092       //modifying particular value

	fmt.Println(c["Pin"])

	//check if key/value exists
	var d = map[string]int{"one": 1, "two": 2, "three": 3}
	val1, ok1 := d["four"] //ok1 will be false since val1 does not find four in the map
	fmt.Println(val1, ok1)
	val2, ok2 := d["three"] //will be true
	fmt.Println(val2, ok2)
	_, ok3 := d["two"] //check if key exists but value is not needed to be stored
	fmt.Println(ok3)

	e := d //e refers to d, so if we change value in d, value will change in e as well
	fmt.Println(e)
	d["two"] = 22 // changing d, but printing e
	fmt.Println(e)
}
