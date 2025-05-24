# In Go language, the program is saved with .go extension and it is a UTF-8 text file.
```go
package main

import "fmt"

func main(){
	fmt.Println("Hello World!")
}
```

- **package main** — Every Go application starts with a main package.
  - Every Go file belongs to a package. —→ Must Needed
  - **main** is a special package — it tells the Go compiler that this file can be run as an executable program.
  - If you don’t use **package main**, you can’t run it directly with go run
  - Think of it as saying: "Hey, Go! This is where my program starts!"
  - main go file er package will be main
  - Every program must start with the package declaration. In Go language, packages are used to organize and reuse the code. In Go, there are two types of program available one is executable and another one is the library. The executable programs are those programs that we can run directly from the terminal and Libraries are the package of programs that we can reuse them in our program.Here, the package main tells the compiler that the package must compile in the executable program rather than a shared library. It is the starting point of the program and also contains the main function in it.

- **import "fmt"** — Imports the fmt package for formatted I/O (like printing).
  - To use built-in or external functionality in Go, you use ``import``.
  - `fmt` is a standard Go package used for formatting text and printing output.
  - `fmt` stands for format, and it allows you to print to the terminal with commands like `fmt.Println()`.
  - Here, `import` keyword is used to import packages in your program and fmt package is used to implement formatted Input/Output with functions.

- **func main()** — Entry point of the Go program.
  - `func` means **function**.
  - `main()` is the **entry point** of any Go executable program.
  - Go starts executing code from inside this function.
  - Without a `main()` function, your program won’t run.
  
- **fmt.Println("Hello, World!")** — Prints the message to the console.
  - `Println()`: This method is present in fmt package and it is used to display `"!... Hello World ...!"` string. Here, Println means Print line.

# How to run Golang Program?
  To run a Go program you need a Go compiler.. Once you have a Go compiler, first you create a program and save your program with extension .go, for example, first.go. Now we run this first.go file in the go compiler using the following command, i.e:
  ```bash
  go run first.go
  ```