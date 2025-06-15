package main

import "fmt"



func main() {
   
	func (a int, b int){
		fmt.Println("Sum: ", a + b)
	}(5, 7)
}
