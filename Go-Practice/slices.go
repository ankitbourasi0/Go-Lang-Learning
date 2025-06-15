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

	fmt.Printf("Slice %v, Length: %d, Capacity of Slice: %d\n", appendElementSlice, len(appendElementSlice), cap(appendElementSlice)) //Slice [1 2 3 4 5 20 21], Length: 7, Capacity of Slice: 10

	//Appending One slice to Another
	slices1 := []int{1, 2, 3}
	slices2 := []int{4, 5, 6}
	slices3 := append(slices1, slices2...)
	fmt.Println(slices3) //[1 2 3 4 5 6]

	//Change the length of a Slice
	arr := [5]int{1, 2, 3, 4, 5}
	Sliced := arr[2:4]
	fmt.Println("Original Slice: ", Sliced)
	Sliced = arr[1:5]
	fmt.Println("New Slice: ", Sliced)

	//💡Memory Efficiency
	//When using slices, Go loads all the underlying element in the memory,
	//If the array is too Large, & You need only few elements.
	//It is better to copy those elements using `copy()` fn.

	//Syntax: copy(destination, source)
	largeSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	fmt.Println("Large Slice: ", largeSlice)
	neededSliceLength := largeSlice[5 : len(largeSlice)-10]
	neededElement := make([]int, len(neededSliceLength))
	fmt.Printf("Needed Slice:%v, %d,  %d\n", neededSliceLength, len(neededSliceLength), cap(neededSliceLength))
	copy(neededElement, neededSliceLength)
	fmt.Printf("New Slice:%v %d %d\n ", neededElement, len(neededElement), cap(neededElement))
}
