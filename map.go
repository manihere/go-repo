/*Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.*/

package main

import("fmt")

func main(){

	//var a = map[keyType]valueType{key1:value1, key2:value2.....}

	var person1 = map[string]int{"Age1": 20, "Salary": 30, "Experience": 10 }

	fmt.Println(person1)

	var Person2 = map[String]string{"Name": "Manish", "Org": "ABC", "lastName": "Panwar"}

	fmt.Println(Person2)

	//create map using make()
	var c = make(map[string]int)
	c["India"] = +91
	c["Pin"] = 110096
	c["Population"] = "1.5B"
	
	fmt.Println("Facts about India \n %v",c)
}