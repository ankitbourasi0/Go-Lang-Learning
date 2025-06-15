package main

import "fmt"

func main() {
	fruits := [4]string{"apple", "banana", "peach", "pear"}
	taste := [2]string{"big", "tasty"}
	fmt.Println("Basic Loop:")
	//Create a basic for loop,
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%s ", fruits[i]) //apple banana peach pear
	}
	fmt.Println("")
	fmt.Println("Order change statement: ")
	//Create a for loop with change in statement order
	for j := 0; j < len(taste); {
		fmt.Printf("%s ", taste[j]) //
		j++
	}
	fmt.Println("")
	fmt.Println("Nested Loop:")
	//Create a nested loop
	for k := 0; k < len(taste); k++ {
		for l := 0; l < len(fruits); l++ {
			fmt.Println(taste[k], fruits[l])
		}
	}
	fmt.Println("")
	fmt.Println("For Loop with RANGE:")
	for idx, value := range fruits {
		fmt.Printf("%d \t %s \n", idx, value)
	}

	fmt.Println("")
	fmt.Println("For Loop with RANGE ommiting index at _:")
	for _, val := range taste {
		fmt.Println(val)
	}
}
