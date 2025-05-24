# In Go (Golang), a data type tells the compiler what kind of data a variable can hold — like numbers, text, or true/false.

![Data Type](assets/datatype.png)

# Types of Go data Types
1. **Bolean types**
   They are boolean types and consists of the two predefined constants: **(a) true (b) false**

2. **Numeric types**
   They are again arithmetic types and they represents a) integer types or b) floating point values throughout the program.

3. **String types**
   A string type represents the set of string values. Its value is a sequence of bytes. Strings are immutable types that is once created, it is not possible to change the contents of a string. The predeclared string type is string.
   ```go
   str := "Hello"
   str[0] = 'M' // ❌ This will give a compile-time error
   You cannot modify str[0] because Go strings are immutable.
   ```

4. **Derived types**
   They include (a) Pointer types, (b) Array types, (c) Structure types, (d) Union types and (e) Function types (f) Slice types (g) Interface types (h) Map types (i) Channel Types

**Array types and structure types are collectively referred to as aggregate types. The type of a function specifies the set of all functions with the same parameter and result types.**

# Boolean Types
In Go, bool is the single type representing Boolean values. It is widely used in logical and conditional operations. It can accept only two values which are "true" or "false".
```go
package main

import "fmt"

func main() {
    // Declare boolean variables
    var var1 bool = true
    var var2 bool = false

    // Print boolean values
    fmt.Println("var1:", var1)
    fmt.Println("var2:", var2)

    // Use boolean in a conditional statement
    if var1 {
        fmt.Println("It's true.")
    } else {
        fmt.Println("It's false")
    }
}
```
```bash
var1: true
var2: false
It's true.
```
### ***🔒 Default value:**
If you declare a bool without assigning a value, it will be false by default.

```go
var isRunning bool
fmt.Println(isRunning) // false
```