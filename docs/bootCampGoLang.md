## Why was Go Created?

**Python**: Easy to use but interpreted hence Slow

**C/C++** : Complex type system, slow compile time as compiler was designed keeping slower older systems in mind

**Java**  : Complex type system

And all these languages had not been built with Concurrency in mind

## What does Go bring in?

- Strong and **statically typed** language
- Excellent **community**
- **Simplicity** is a key feature
- Fast compile time
- Garbage collection by Go Runtime
- Built-in **Concurrency**
- Compiles to stand alone binaries (**PS** I[KK] love this feature)


## GOROOT, GOPATH, GOOS, GOARCH

- `GOROOT` is not really required env var to setup if you install go in default loction.
- `GOROOT` is always been pointer to root of Go installation. It is similar to `JAVA_HOME` in Java.
- You should not set $GOROOT because the correct value is already embedded in the go tool.
- Setting `$GOROOT` will override the value stored in the go tool which could lead to the `go tool` from one version of Go pointing to the compiler and standard library from another version.
- `GOPATH` is path where you will let GO download all libraries it needs when you run `go get ...` command. `GOPATH` is a compound path, so you could do following:
```bash
export GOPATH=/home/kkulkarni/golib # this is where `go get` command downloads data
export PATH:$PATH:$GOPATH # this is because GOPATH would contain `bin` that you may need code from

# But if you want to keep your workspace separated from libs, you can create additional directories and do this:
export GOPATH=$GOPATH:/home/kkulkarni/go-workspace
export PATH=$PATH:$GOPATH
```
With above 2 being in your GOPATH, 1st path will be where all libs will go but 2nd path is where you will put your apps that you are working on.

- GOOS and GOARCH became totally optional following GO 1.0 release. They are only required when cross-compiling.

## Directory Structure

Within golang you'd need to have following dirs in your workspace:

- **bin** : this is where any binary will be installed if you run `go install`
- **src** : this is where you'd put your source code. It requires you to put src code in nested directories matching path of your github or source code repo. So if this repo is where I am developing code, the my `src` dir should contain `src/github.com/kedark3/its-go-time`
- **pkg** : this is where you may find intermediate binaries built by go during compile time


## Ways to run your go application

- `go run` will compile and run your app directly without generating new binaries in your `pwd`
- `go build` will compile and generate an executable binary in your `pwd`
- `go install` will compile and generate an executable binary in your `bin` dir within your workspace



## Variables

### Variable Declaration

**Option1:**
```go
    var <varname> <type>
    // e.g.
    var i int
    var s string
```

**Option2:**
```go
    var <varname> <type> = <value>
    // e.g.
    var i int = 42
    var s string = "abcd"
```

**Option3:**
```go
    <varname> := <value>
    //e.g.
    i := 42
    s := "Hello world"
```


**Option4:**
```go
    var(
        i int = 1
        j string = "abcd"
    )
```


**Notes**: 
- If a variable is declared at package level, you can `shadow` it with new declaration at local level within a function.
- You **cannot** have any variables unused as compiler would complain about it.


### Variable Naming

- `lower case` variable is only **scoped within package**
- `upper case` letter in the beginning is **visable outside of package**
- Variables within functions or blocks are **block scoped**
- **Single Letter** variable names are good for very short lived variable like in loops
- **Descriptive** names are encouraged for variables used across the packages
- **Acronyms** should be all uppercase e.g. `theUrl` is bad, better is`theURL` 


### Variable Type Conversions

- Conversion has to be **explicit**
```go

    var i int = 42
    var j float64
    j = float64(42)
    j = i // this won't work as go compiler won't risk implicitly converting types and losing data
    // e.g. if you had
    var k float64 = 42.6
    var l int = k // if this was to work in Go, it would've been converted to 42, losing info about 0.6
```


- If you need to convert **int** and **string** recommended to use `strcov` package otherwise `string(42)` may make result in a `*`(unicode) character vs `"42"` which you intended to do.


## Primitives

### Boolean Type

- Simplest primitive type.
- leaving uninitialized var in Golang means its initialize to 0 value, so for boolean it is `false`

```go
    var m bool = true  // or false
    var n bool // leaving uninitialized var in Golang means its initialize to 0 value, so for boolean it is `false`
    fmt.Println("%v, %T",n , n) // %v for value and %T for type
```

### Integer Type

- an `int` can be 32/64 bits or sometimes upto `128` bits depending on the platform you are using. You can have `8-bit`, `16-bit` as well.
- an `unint` can be 8/16/32 bits only, we don't have 64 bit option.
- a `byte` is an alias for `uint8` as most datastreams are encodeded in `bytes`
- **Arithmatic operations**:
```go
    a := 10
    b := 3

    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Println(a / b) // int result 10/3 = 3 and it drops remainder
    fmt.Println(a % b) // remainder
```

- You cannot operate on different types

- **bit operators**:
```go
    a := 10 // 1010
    b := 3  // 0011

    fmt.Println(a & b) // AND 0010
    fmt.Println(a | b) // OR  1011
    fmt.Println(a ^ b) // XOR 1001
    fmt.Println(a &^ b) // AND NOT 1000
    a = 8 // 2 ^ 3
    fmt.Println(a << 3) // 2 ^3 * 2^3 = 2^ 6 = 64
    fmt.Println(a >> 3) // 2 ^3 / 2 ^3 = 2^0 = 1
```



### String Types

- `String` is any UTF-8 char in golang.
- `String`s in Go are aliases for byte
- `String`s are immutable

```go
package main

import (
    "fmt"
)

func main(){
    s := "this is a string"
    fmt.Printf("%v, %T", s, s) // prints "this is a string, string"
    fmt.Printf("%v, %T", s[2], s[2]) //  prints 105, uint8
    fmt.Printf("%v, %T", string(s[2]), s[2]) //  prints "i", uint8

    s[2] = "u" // this is NOT VALID

    s2 := "Another string"
    fmt.Println(s+s2) // concats the strings

    b := []byte(s) // converting string to slice of bytes
    fmt.Println("%v, %T",b , b) // prints [] of ascii values of each character and type is []uint8
    s := "this is a string"
	b := []byte(s)
	fmt.Println(string(b))
	b[1] = 105 // 'i'
	b[2] = 97 // 'a'
	fmt.Println(string(b)) // modifies original string to be "tias is a string"
}
```


- `rune` represents any UTF-32 character, it is an alias for int32
```go
    r := 'a' // single quotes is a rune
```



# Constants

- begins with keyword `const`
- Uses camelCase or first letter caps if you want to export it
- cannot be set to something that can't be evaluated at compile time and needs execution
- cannot be of collection types such as arrays, maps
- can shadow/override const from package level in function level - but not recommended as could be confusing
- Untyped constants like `const a =42` acts as literal, and it will simply be replacing all occurences of `a` in the program with literal `42` and it can interoperate with similar types, so in this example it will work with int8/16/32/64

## Enumerated constants

```go
package main 

import "fmt"

const(
    a = iota
    b = iota
    c = iota
)

const(
    // _ = iota in case we don't care about 0 value so d,e,f will be 1,2,3
    d = iota // value resets to 0 again in this block
    e // compiler assumes it to be =iota
    f // compiler assumes it to be =iota
)

const(
    x = iota + 5 // arithmatic operations is allowed
    y
    z
)
func main(){
    fmt.Printf("%v,\n",a)  // 0
    fmt.Printf("%v,\n",b)  // 1
    fmt.Printf("%v,\n",c)  // 2

    fmt.Printf("%v,\n",d)  // 0
    fmt.Printf("%v,\n",e)  // 1
    fmt.Printf("%v,\n",f)  // 2
}
```


# Arrays and Slices

## Arrays

- Declaration:

Syntax 1:
```go
    arr := [3]int{1,2,3}
```
Syntax 2:
```go
    arr := [...]int{1,2,3}
```
Syntax 3:
```go
    var arr [3]strings
    arr[0]= "a"
```

- They are allocated contiguous memory block
- Should always have homogeneous type
- Nested array, e.g. Matrix:
```go
    var identityMatrix [3][3]int = [3][3]int{[3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}

```
OR
```go
    var identityMatrix [3][3]int 
    identityMatrix[0]= [3]int{1,0,0}
    identityMatrix[1]= [3]int{0,1,0}
    identityMatrix[2]= [3]int{0,0,1}

```

- Arrays are actually considered values, not reference, and passing big array to functions can cause program to slow down as it creates copy
```go
    a := [...]int{1,2,3}
    b := a
    b[1] = 5

    fmt.Println(a) // [1,2,3]
    fmt.Println(b) // [1,5,3]
```

- Arrays are limited in use because of static length

## Slices

- It is backed by an Array
```go
    a:= []int{1,2,3}
```
- Has `length` and `capacity` functions `len()` and `cap()`
- These are reference type
```go
    a := []int{1,2,3} // defines a slice
    b := a
    b[1] = 5 // change is made in `a` as well because slices are reference type

    fmt.Println(a) // [1,5,3]
    fmt.Println(b) // [1,5,3]
```

- Different ways to declare
```go
    a := []int{1,2,3,4,5,6,7,8,9,10} // this is slice but could have been [...]int array
    b := a[:] // slices of all elements
    c := a[3:] // slice from 4th element to end
    d := a[:6] // first 6 elements, read this as [:6) where element at 6th index is not included
    e := a[3:6] // slice the 4th,5th,6th elements, read this as [3:6) where element at 6th index is not included
    a[5] = 100 // changes all slices
```
- Using `make` function
Syntax 1:
```go
    a := make([]int, 3) // (type, length) created [0,0,0], capacity set to 3 as default matching length
```
Syntax 2:
```go
    a := make([]int, 3, 100) // (type, length)created [0,0,0], capacity cap(a) set to 100 though, this is good way to avoid re-copying elements from smaller array to larger array as items in the slice grows, because slices usually point to an array and when we need to add more elements than length of array, slice does copy operation from smaller array to larger array and increases capcity automatically
```

- Adding elements to a slice:
```go
    a := []int{} // len(a) 0 and cap(a) 0
    a = append(a, 1) // len(a)=1 and cap(a)=2 - adding new element caused go to create a new array with larger capcity, copying elements from small array to large array and that's why cap(a) 2

    a = append(a, 2, 3, 4, 5) // cap(a)=8 now, Golang may just keep doubling array sizes(hence capcity) but that is upto go runtime to decide

    a = append(a, []int{1,2,3,4}) // will NOT WORK as you can't append slices like this
    a = append(a, []int{1,2,3,4}...) // this will WORK as `...` causes slices to be decomposed to append(a, 1,2,3,4) `...` is called a spread operator

    a := []int{1,2,3,4,5}
    b := a[1:] // drops 0th element
    c := a[:len(a)-1] // drops last element
    d := append(a[:2], a[3:]...)
```

# Maps 

- Map has [keys]values of [type1]type2
- Map keys have to be able to be tested for equality, so map key can't be a slice

# Structs



Based on [Free GOlang Bootcamp](https://www.youtube.com/watch?v=YS4e4q9oBaU*)

