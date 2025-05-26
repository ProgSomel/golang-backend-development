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

  - To use built-in or external functionality in Go, you use `import`.
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

**----------------------------------------------------------------------------------------------------------------------------------**

# Go Compiler and Runtime

## If go compiler is changing the code to machine code, won’t the machine understand the command code directly? Why there is need for go runtime?

## Go Compiler

Go is indeed a compiled language that's means go convert program to **machine code** specific to target platform.
If i'm making the code, if i'm building the executable on an x86 processor, then the target platform of that executable is x86 architecture.
If i'm making the code, if i'm building the executable on an ARM architecture, then the executable will be able to run on ARM based processor.

so yes, the go source code is translated into machine code, and when we are doing that, we are making the machine code specific to a target platform and this is done by go Compiler.

This machine code is what the computer's processor can execute directly.

## Go Runtime

Despite compiling to machine code, go programs require a Runtime environment to manage certain aspects. Those related to memory management, which is garbage collection and also concurrency and along with that, there are certain other language features that the go Runtime handles. First thing id garbage collection. Go used automatic garbage collection to manage memory allocation and deallocation, this means that the Runtime is responsible for identifying and reclaiming memory that is no longer in use, which helps developers avoid memory leaks and manage memory efficiently.
Now come to the Go concurrency and Goroutine. Goes concurrency model relies heavily on go routine. Lightweight threads managed by go runtime. The runtime scheduler, it is part of the runtime efficiently. Schedules go routines across available CPU cores, allowing for concurrent execution without developers needing to manage threads manually.
And when we are talking about go language, we mean the go language, the go syntax, the design of the language along with the go compiler and go runtime. And this complete package of go language helps us make better programs, helps us make highly efficient programs
Go programs also depend on runtime libraries provided by the go runtime, such as those for reflection, channel operations and other language features. These libraries help facilite common programming tasks and provide support for go specific features. The go runtime also ensures that go programs behave consistently across different operating systems and hardware architectures, abstracting away platform specific details and providing a uniform execution environment
**In summary, while Go compiles to machine code, the go runtime is essential for managing memory, supporting concurrent execution via go routines, providing runtime libraries, and ensuring cross-platform compatibility.**