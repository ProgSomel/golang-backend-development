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

**--------------------------------------------------------------------------------------------------------------------**

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

**--------------------------------------------------------------------------------------------------------------------**

# Numeric Types
## integer
```md
| Sr.No. | Type    | Description                                   |
|--------|---------|-----------------------------------------------|
| 1      | `uint8`  | Unsigned 8-bit integers (0 to 255)            |
| 2      | `uint16` | Unsigned 16-bit integers (0 to 65535)         |
| 3      | `uint32` | Unsigned 32-bit integers (0 to 4294967295)    |
| 4      | `uint64` | Unsigned 64-bit integers (0 to 18446744073709551615) |
| 5      | `int8`   | Signed 8-bit integers (-128 to 127)           |
| 6      | `int16`  | Signed 16-bit integers (-32768 to 32767)      |
| 7      | `int32`  | Signed 32-bit integers (-2147483648 to 2147483647) |
| 8      | `int64`  | Signed 64-bit integers (-9223372036854775808 to 9223372036854775807) |
```
**int**	Depends on system (32/64 bits)
**uint**	System-dependent
**uint**	System-dependent

**Example**
The following is an example of how the integer types are used in Go:
```go
package main

import "fmt"

func main() {
    // Unsigned integers
    var u8 uint8 = 255
    var u16 uint16 = 65535

    // Signed integers
    var i8 int8 = -128
    var i16 int16 = 32767

    // Printing values
    fmt.Println("Unsigned 8-bit integer (uint8):", u8)
    fmt.Println("Unsigned 16-bit integer (uint16):", u16)
    fmt.Println("Signed 8-bit integer (int8):", i8)
    fmt.Println("Signed 16-bit integer (int16):", i16)
}
```
```bash
Unsigned 8-bit integer (uint8): 255
Unsigned 16-bit integer (uint16): 65535
Signed 8-bit integer (int8): -128
Signed 16-bit integer (int16): 32767
```
## ⚠️ Important Notes:
- Use **signed types** when you expect negative values.
- Use **unsigned types** if you know the value can't be negative — e.g., file sizes, memory addresses.
- Be cautious mixing **signed and unsigned types** in expressions.

**--------------------------------------------------------------------------------------------------------------------**

## Floating-Point and Complex Types

| Sr.No. | Type        | Description                                                  |
|--------|-------------|--------------------------------------------------------------|
| 1      | `float32`    | IEEE-754 32-bit floating-point numbers                       |
| 2      | `float64`    | IEEE-754 64-bit floating-point numbers                       |
| 3      | `complex64`  | Complex numbers with `float32` real and imaginary parts      |
| 4      | `complex128` | Complex numbers with `float64` real and imaginary parts      |

> **Note:** The value of an *n*-bit integer is *n* bits and is represented using two's complement arithmetic operations.
> 🔍 **float64** is the default for floating-point numbers in Go.

**Example**
The following is an example of how the float types are used in Go:
```go
package main

import "fmt"

func main() {
    var pi float32 = 3.14159
    var e float64 = 2.718281828459045

    var firstComplexNumber complex64 = complex(1.0, 2.0)
    var secondComplexNumber complex128 = complex(3.0, 4.0)

    // Printing values
    fmt.Println("Value of pi (float32):", pi)
    fmt.Println("Value of e (float64):", e)
    fmt.Println("First complex number (complex64):", firstComplexNumber)
    fmt.Println("Second complex number (complex128):", secondComplexNumber)

    // Printing real and imaginary parts
    fmt.Println("Real part of first complex number:", real(firstComplexNumber))
    fmt.Println("Imaginary part of first complex number:", imag(firstComplexNumber))

    fmt.Println("Real part of second complex number:", real(secondComplexNumber))
    fmt.Println("Imaginary part of second complex number:", imag(secondComplexNumber))
}
```
```bash
Value of pi (float32): 3.14159
Value of e (float64): 2.718281828459045
First complex number (complex64): (1+2i)
Second complex number (complex128): (3+4i)
Real part of first complex number: 1
Imaginary part of first complex number: 2
Real part of second complex number: 3
Imaginary part of second complex number: 4
```
**⚠️ Key Points:**
- Use **float32** when memory is a concern.
- Prefer **float64** for most accurate calculations.
- Be aware of floating-point **rounding errors** (they're not always 100% precise).