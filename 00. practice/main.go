package main

import "fmt"



func main() {
    ptr := new(int);
	*ptr = 20;

	fmt.Println(*ptr)
}
