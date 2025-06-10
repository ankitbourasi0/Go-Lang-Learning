package main

import (
	"fmt"
)

func main() {
	//1.Create an Empty array using var
	var myarray = [5]int{}
	fmt.Println("Empty Integer Array: ", myarray) // [0 0 0 0 0]
	//💡default values of int is 0, string is "", bool is false, float is 0

	//2.Create an Empty array using shortcut
	myStringArray := [4]string{}
	fmt.Println("Empty String Array: ", myStringArray) //[   ]

	//3.Create a bool array using shortcut
	myboolArray := [3]bool{}
	fmt.Println("Empty Boolean Array: ", myboolArray) // [false false false]

	//4.Create a float array using var
	var myFloatArray = [3]float32{}
	fmt.Println("Empty Float Array: ", myFloatArray) //[0 0 0]

	//5.Create an empty Inferred array
	myInferredArray := [...]string{}
	fmt.Println("Empty Inferred Array: ", myInferredArray) //Empty Inferred Array:  []

	//6.Create an Inferred Array of Some element
	myInferredArray2 := [...]string{"A", "B", "C"}
	fmt.Println("String Inferred Array: ", myInferredArray2) //String Inferred Array:  [A B C]

	//7.Create an Array using var of Size 5
	var myarray2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myarray2) //[1 2 3 4 5]

	//8.Crate an Array with Intializing Specific element
	var myarray3 = [5]int{2: 20, 4: 40}
	fmt.Println(myarray3) // [0 0 20 0 40]

	//9.Get len of Intialize Specfici element Array
	fmt.Println("Length: ", len(myarray3)) //Length:  5

	//10.Access a Element of an Array
	myNewArray := [...]string{"Volvo", "Suzuki", "MG"}
	fmt.Printf("3rd car is: %v \n", myNewArray[2]) //3rd car is: MG

	//11.Change a Element of the array
	myNewArray[1] = "Gopher"
	fmt.Println(myNewArray) //[Volvo Gopher MG]

	//12. Create a unintialized Array of type int
	var myUninitalizedArray []int
	fmt.Println(myUninitalizedArray) //[]
}
