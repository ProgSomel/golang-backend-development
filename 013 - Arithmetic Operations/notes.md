# Arithmetic Operations

```go
package main

import "fmt"

func main(){
	//? variable declaration
	var a, b int = 10, 5

	var result int

	result = a + b;
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const pi float64 = 22/7;
	fmt.Println(pi)



}
```
```bash
Addition:  15
Substraction:  5
Multiplication:  50
Division:  2
Remainder:  0
3
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//? variable declaration
	var a, b int = 10, 5

	var result int

	result = a + b;
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const pi float64 = 22/7.0;
	fmt.Println(pi)



}
```
```bash
Addition:  15
Substraction:  5
Multiplication:  50
Division:  2
Remainder:  0
3.142857142857143
```

------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//? variable declaration
	var a, b int = 10, 5

	var result int

	result = a + b;
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const pi int = 22/7.0;
	fmt.Println(pi)



}
```
```bash
cannot use 22 / 7.0 (untyped float constant 3.14286) as int value in constant declaration (truncated)
```

**--------------------------------------------------------------------------------------------------------------------------------------------**

# Overflow and underflow
