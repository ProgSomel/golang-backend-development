# ✅ Go (Golang) — Hello, World Program:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
- 🧠 Explanation:
  ```go
  Line	Description
  package main	Every Go program starts with a package declaration. main is the entry point of an executable program.
  import "fmt"	This imports the fmt package, which is used for formatted I/O (like printing to the console).
  func main()	The main function — the starting point of execution.
  fmt.Println(...)	Prints the string to the console with a newline at the end.
  ```

- 🛠️ How to Run:
  You can run this using:
  [Go Playground](https://go.dev/play/)
  Your local Go setup using **go run filename.go**

# Go Compiler
- Compiler: translator for your code(translate to machine code which is binary)
- Runtime: manages the execution of your program, memory allocation, garbage collection(cleaning up unused memory) etc