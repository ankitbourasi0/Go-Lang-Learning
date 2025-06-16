package main

import "fmt"

func main() {
	//A map is an unordered and changeable collection, that does not allow duplicates
	//Default value of a map is nil.
	//Map holds references to an underlying hashtable.

	//Create a empty map using var & :=
	var myMap map[string]string
	var myMap2 = map[string]string{}
	myMap3 := map[string]string{}
	fmt.Println(myMap, myMap2, myMap3) //map[] map[] map[]

	//Create map with some values using var & :=
	var myMap4 = map[int]string{1: "Apple", 2: "Orange", 3: "Grapes"}
	fmt.Println(myMap4) //map[1:Apple 2:Orange 3:Grapes]

	//order of map element is might different from code to output cuz,
	//The map is stored element in efficient data retrieval form.

	//add a value in map
	myMap2["Brand"] = "BMW"
	myMap2["Car"] = "Forex"
	fmt.Println(myMap2) //map[Brand:BMW Car:Forex]

	//Change a value in Map
	myMap2["Car"] = "X6 Coupe SUV"
	fmt.Println(myMap2) //map[Brand:BMW Car:X6 Coupe SUV]

	//Creat a map using Make() fn
	var mapusingMake = make(map[int]string)
	var mapusingMake2 = make(map[int]string)
	fmt.Println(mapusingMake, mapusingMake2) //map[] map[]

	//The right way to Create an empty map is using Make() fn ,
	//cuz if you use other way might it cause runtime error
	var a = make(map[int]string)
	var b map[int]string
	fmt.Println(a == nil) //false
	fmt.Println(b == nil) //true

	//delete an element from map
	delete(myMap4, 2)   //delete key 2 and its value
	fmt.Println(myMap4) //map[1:Apple 3:Grapes]

	//check for specific element in a map
	//check existing key nd its value
	val, ok := myMap4[1]
	fmt.Println("Val: ", val, "Key: ", ok) //Val:  Apple Key:  true
	//you can use blank identifier(_) in a place of val to check only key
	_, k := myMap4[1]
	fmt.Println("Key: ", k) //Key:  true

	//check non-existing key, 💡argument key should be always of type that key is using the map.
	_, key := myMap4[4]
	fmt.Println("Key: ", key) //Key:  false

	//Map references
	//if 2 maps refers to same hash table, change in one affect another.
	c := map[int]string{1: "Apple", 2: "Linux", 3: "Hp"}
	d := c
	fmt.Println(c) //map[1:Apple 2:Linux 3:Hp]
	fmt.Println(d) //map[1:Apple 2:Linux 3:Hp]
	d[2] = "Asus"
	fmt.Println("Change after c")
	fmt.Println(c) //map[1:Apple 2:Asus 3:Hp]
	fmt.Println(d) //map[1:Apple 2:Asus 3:Hp]

	//Iterating over Maps
	e := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range e {
		fmt.Printf("%v : %v,", k, v)
	}
	fmt.Println(e) //map[four:4 one:1 three:3 two:2]

	//Iterating over Maps in specific Order.
	var order = []string{"one", "two", "three", "four"}
	for _, element := range order {
		fmt.Printf("%v : %v ", element, e[element]) //one : 1 two : 2 three : 3 four : 4
	}
}
