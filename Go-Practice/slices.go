package main

import "fmt"

func main() {
	//Create a Empty Slice
	mySlice := []int{}
	fmt.Println("Empty Slice: ", mySlice) //[]

	//Create a Empty Slice using Make Function
	var mySliceUsingMake = make([]int, 0)                     //Syntax: make([]datatype, length, capacity) 💡(Capacity is optional)
	fmt.Println("Empty Slice using Make: ", mySliceUsingMake) // []

	//Create a Slice of type String
	var myStringSlice = []string{"Apple", "Bananas", "Pear"}
	fmt.Println("String Slice : ", myStringSlice) //[Apple Bananas Pear]

	//Create a Intger Slice
	myIntSlice := []int{1, 2, 3, 4}
	fmt.Println("Int Slice : ", myIntSlice) //[1 2 3 4 ]

	//Create Slice from An Array
	mySliceFromArray := []int{}
	myArray := [5]int{1, 2, 3, 4, 5}
	mySliceFromArray = myArray[2:4]                      //Start from Index 2 upto 4th index , 4th index is excluded
	fmt.Println("Slice from Array : ", mySliceFromArray) // [3 4]

	//Create a Empty Slice of length 5 and capacity 10
	myLengthSlice := make([]int, 5, 10)
	fmt.Printf("Array: %v, Length:%d, Capacity:%d \n", myLengthSlice, len(myLengthSlice), cap(myLengthSlice)) // [0 0 0 0 0], Length:5, Capacity:10

	//Create a Slice of Same length and Capacity
	var myEqualSlice = make([]int, 5)
	fmt.Printf("Equal Slice: %v,  Length:%d, Capacity:%d  \n", myEqualSlice, len(myLengthSlice), cap(myEqualSlice)) //[0 0 0 0 0],  Length:5, Capacity:5

	//Append few element in a Slice
	appendElementSlice := []int{1, 2, 3, 4, 5}
	appendElementSlice = append(appendElementSlice, 20, 21)
	fmt.Println(appendElementSlice) // [1 2 3 4 5 20 21]

	/*Check its Capacity, 💡whenever slice is appending a element and its exceeds the length its capacity is double
	via underlying array, It grows Until element reached 1024, after this capacity increased by 1.25x
	*/

	fmt.Printf("Slice %v, Length: %d, Capacity of Slice: %d", appendElementSlice, len(appendElementSlice), cap(appendElementSlice)) //Slice [1 2 3 4 5 20 21], Length: 7, Capacity of Slice: 10
}
