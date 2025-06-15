package main

import "fmt"

func main(){
	//? variable declaration
	// var a, b int = 10, 5

	// var result int

	// result = a + b;
	// fmt.Println("Addition: ", result)

	// result = a - b
	// fmt.Println("Substraction: ", result)

	// result = a * b
	// fmt.Println("Multiplication: ", result)

	// result = a / b
	// fmt.Println("Division: ", result)

	// result = a % b
	// fmt.Println("Remainder: ", result)

	// const pi  = 22/7.0;
	// fmt.Println(pi)

	//? overflow with signed integers
	var maxInt int64 = 2222223777777722222
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

}