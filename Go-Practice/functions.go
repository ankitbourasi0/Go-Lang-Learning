package main

import "fmt"

func main() {
	//create a function for greeting and call it.
	greet("Ankit") //Hello Ankit

	//create a function with return type for int.
	fmt.Println("My age is", age(23)) //My age is 23

	//create a function to return multiple parameters
	var name, age, job = introduce("Ankit", 23, "Software Engineer")
	fmt.Printf("My name is %s, I'm %d years old and I worked as a %s\n", name, age, job)
	//My name is Ankit, I'm 23 years old and I worked as a Software Engineer

	//create a function with named return value.
	fmt.Println(add(2, 3)) //5

	//create a function with multiple name return value and store them in variables
	var product, sum = calc(2, 3)
	fmt.Printf("Product is %d & Sum is %d\n", product, sum) //Product is 5 & Sum is 6

	//create a recursion fn base condition of 5
	testCount(1) //1 2 3 4

	fmt.Println("\nFactorial of 4 is:", factorial(4)) // Factorial of 4 is: 24
}

func greet(name string) {
	fmt.Printf("Hello %s\n", name)
}

func age(age int) int {
	return age
}

func introduce(name string, age int, job string) (string, int, string) {
	return name, age, job
}

func add(a, b int) (result int) {
	result = a + b
	return
}

func calc(a, b int) (product int, sum int) {
	product = a + b
	sum = a * b
	return product, sum
}

func testCount(count int) int {
	if count == 5 {
		return 0
	}
	fmt.Print(count, " ")
	return testCount(count + 1)
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
