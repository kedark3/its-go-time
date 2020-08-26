# Basic Rules of Go

- Go is compiled language
- It has similarity in syntax with C, C++, C# as well as Pascal family of languages
- Initially introduced in 10 November 2009
- Case sensitive
- Tries to reduce the code a developer has to type
- Uses braces for code blocks
- No semicolons required but lexer might add those for you at end of the statements
- Go uses `mixedCase` for naming items
- If a given name starts with Capital letter - mixed case, it is equivalent to public declaration in C++ or Java
- There is no concept of classes but there are lot of structs and types and interfaces
- Requires no VM on your machine for Go to execute, as compiled code will include required runtime, hence the compiled object code may be much larger than source code as it statically links runtime into it
- Compiled code will execute on specific platform

# Basic Data types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type. 

Printing value and type of a variable:
```fmt.Printf("Type: %T value: %v\n", var_name, var_name```


# Variables

Variables declared without an explicit initial value are given their zero value.

The zero value is:

    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.

Inside a function, the `:= `short assignment statement can be used in place of a var declaration with implicit type. 

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

# Type conversions

 The expression T(v) converts the value v to the type T.

Some numeric conversions:
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
Or, put more simply:
```
i := 42
f := float64(i)
u := uint(f)
```
Unlike in C, in Go assignment between items of different type requires an explicit conversion. 



# Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

```
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

