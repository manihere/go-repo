package main

import (
	"fmt"
)

/*func main() {
for i := 0; i < 5; i++ {
	if i == 3 {
		continue
	}
	fmt.Println(i)
}*/

// continue is used to skip a particular iteration without breaking code
func main() {

	for i := 0; i < 5; i += 1 {
		if i == 3 {
			continue //if i==3, skip this block and continue loop
		}
		fmt.Println("Using continue function, i= ", i)
	}

	//break is used to terminate loop
	for i := 1; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println("Using break function, i= ", i)
	}

	//nested loops
	adj := []string{"big", "small"}
	fruits := []string{"apple", "banana", "pear"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	//range is used to iterate easily in array
	//range returns both index and value
	//index(idx) prints index pos value, val(value) provides value
	veg := []string{"carrot", "potato", "onion"}
	for idx, val := range veg {
		fmt.Printf("\n Printing index and value: %v %v", idx, val)
	}
	//use _ to skip calculating either index or value
	for idx, _ := range veg {
		fmt.Printf("\n Printing only index: %v", idx)
	}

	for _, val := range veg {
		fmt.Printf("\n Printing only value: %v", val)
	}
}
