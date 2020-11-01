Based on [Free GoLang Bootcamp](https://www.youtube.com/watch?v=YS4e4q9oBaU*)
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

- Declaration syntax:
```go
map1 := make(map[string]int) // use this if you don't want to add values to map immediately

// Alternative syntax

map1 := map[string]int{
    "a":1,
    "b":2,
    "c":3,
}

```
- Accessing or modifying values:
```go

fmt.Println(map1["b"])

map1["d"] = 4
```
- **NOTE:** Iterating over Map is not ordered, so reading map does not guarantee any order.
- Deleting values from map:
```go
delete(map1, "c")

fmt.Println(map1["c"]) // now it prints 0 even if the "c" is deleted, which is incorrect

// Safer way is

pop, ok := map1["c"]

fmt.Println(map1["c"], ok) // this tells us with ok=false and tells us key is missing
```

- Checking length:
```go
fmt.Println(len(map1)) // prints length i.e. number of items in map
```

- Passing maps around:

```go
// if you do this
map2 := map1

delete(map2, "a")

//then print both
fmt.Println(map1)
fmt.Println(map2) // both of them will be missing key "a" as map2 just points to map1

```

# Structs

- Declaration:
```go
// Dr. Who
type Doctor struct{
    number int
    actorName string
    companions []string
}
```

- Initialization:

```go

func main(){
    aDoctor := Doctor{
        number: 3,
        actorName: "John Pertwee",
        companions: []string{
            "Liz",
            "Jo",
            "Sarah",
        }
    }

    fmt.Println(aDoctor)
    fmt.Println(aDoctor.actorName)
    fmt.Println(aDoctor.companions[1]) // extract item from Slice

}
```

- You can also ignore a field from struct by omitting it:

```go
type Doctor struct{
    number int                 // If you'd want to let other programs see this struct and its members
    actorName string           // You'd want to name the struct and its members with Capital letter beginning
    episodes []string          // In this case, only struct is visible but not its items
    companions []string
}

func main(){
    aDoctor := Doctor{
        number: 3,
        actorName: "John Pertwee",
        companions: []string{  // Here I am skipping episodes and its fine
            "Liz",
            "Jo",
            "Sarah",
        }
    }

    fmt.Println(aDoctor)
    fmt.Println(aDoctor.actorName)
    fmt.Println(aDoctor.companions[1]) // extract item from Slice

}
```
-  If you'd want to let other programs see this struct and its members,  You'd want to name the struct and its members with Capital letter beginning.

- Anonymous Struct, which can be used in web communication for accepting return data and forwarding it to another program,
  just create a anonymous struct:

```go

func main(){
    aDoctor := struct{name string}{name: "John Pertwee"}
    fmt.Println(aDoctor)
    fmt.Println(aDoctor.name)
}

```

- **Struct Copy** is independent. Unlilke, slices, arrays and maps.

```go

func main(){
    aDoctor := struct{name string}{name: "John Pertwee"}
    fmt.Println(aDoctor)
    fmt.Println(aDoctor.name)
    anotherDoctor := aDoctor // creates copy
    anotherDoctor.name = "Tom Baker"
    fmt.Println(aDoctor.name) // Prints John Pertwee
    fmt.Println(anotherDoctor.name) // Prints Tom Baker
}
```

- You could pass same struct to another name using Pointers and References:

```go
func main(){
    aDoctor := struct{name string}{name: "John Pertwee"}
    fmt.Println(aDoctor)
    fmt.Println(aDoctor.name)
    anotherDoctor := &aDoctor // creates pointer to aDoctor
    anotherDoctor.name = "Tom Baker"
    fmt.Println(aDoctor.name) // Prints Tom Baker
    fmt.Println(anotherDoctor.name) // Prints Tom Baker
}
```

- Embedding (sort of like Inheritance but does not provide Polymorphism, you can't use one object interchangeably with another)

Syntax 1:
```go
type Animal struct{
    Name string
    Origin String
}

type Bird struct{
    Animal // embedding Animal, it creates "Has A" relation, "Bird Has A Animal" but not "Is a Animal" like traditional inheritance
    SpeedKPH float32
    CanFly bool
}

func main(){
    b := Bird{}
    b.Name = "Emu"
    b.Origin = "Australia"
    b.SpeedKPH = 48
    b.CanFly = false
    fmt.Println(b.Name)
}

```
Syntax 2:
```go

b := Bird{
    Animal: Animal{Name: "Emu", Origin: "Australia"},
    SpeedKPH: 48,
    CanFly: false,
}
```


- Tags:

```go
tpye Animal struct{
    Name string `required max: "100"` // makes this required with max length "100"
    Origin string
}

func main(){
    t := reflect.TypeOf(Animal{}) // need to import "reflect" pkg
    field, _ := t.FieldByName("Name") // extract field from Animal
    fmt.Println(field.Tag) // This prints tag 'required max:"100"'
}
```



# Control Flow

## If statements

- Basic Syntax:

```go

if <expression>{

}else if <expression{

}else{

}

```

- Initializer Syntax:
```go
if pop, ok := map1["key"]; ok{ // we initialize pop, ok and use ok as expression
    fmt.Println(pop) // pop is also accessible in this scope
}
```

- Comparison Operators:
```go
==
<
>
<=
>=
!=
```
- Logical Operators:
```go
|| // OR
&& // AND
! // NOT
```

- Floating point or decimal value comparison is tricky as Go has approximation of floating point numbers and not absolute. Hence `==` operator with floating point number is is not a good idea.




## Switch Statements


- Syntax:
```go

func main(){
    switch <tag>{
        case <1>:
        // optionally break  if data doesn't match expectations and you want to exit out of switch before doing remaining `something` below in this particular case
        // do Something
        case <2>:
        // do Something
        default:
        // do Something
    }
}

```
- Syntax2:
```go
func main(){
    switch 3{
        case 1, 5, 10: // useful to compare multiple cases
        // do Something
        case 2, 4, 6:
        // do Something
        default:
        // do Something
    }
}

```

- Syntax 3:

```go
func main(){
    switch i:= 2+3, i{
        case 1, 5, 10:
        // do Something
        case 2, 4, 6:
        // do Something
        default:
        // do Something
    }
}


```
- Tagless Syntax:

```go
i := 10
switch{

    case i <= 10:
        // do Something
    case i <= 20:
        // do Something
    default:
        // do Something
}
```

- `fallthrough` 

```go
i := 10
switch{

    case i <= 10:
        // do Something
        fallthrough //if this case each matched, it continues going through and executing other cases too
    case i <= 20:
        // do Something
    default:
        // do Something
}
```

- Type based switch:

```go

func main(){
    var i interface{} =1
    switch i.(type){
        case int:
        // do something
        case float64:
        // do something
        case string:
        // do something
        case [3]int: // only matches array with 3 elements of int type
        // do something
        case []int:
        // do something
        default:
        // do something
    }
}

```


# Looping

## for loops

### simple loops

- Basic Syntax:

```go
// option 1
for i := 0; i < 5; i++{

}

// option 2

for i :=0; i< 5; i = i+2{

}

// INCORRECT way to do option 3 
for i :=0, j :=0; i<5; i++, j++{}

// CORRECT way
for i, j := 0, 0; i < 5; i, j = i+1, j+1{} // i,j = i++, j++  -- won't work

```

- You can leave out initialization of variable outside of loop:
```go

i := 1 // this makes `i` scoped to `main()` function
for ; i< 5; i++{  // semicolon is still important

}
fmt.Println(i) // this is valid
```

- We can drop incrementing value, this is like a `while` loop:

```go
i := 0
for ; i<5 ; { // semicolon is still important

    fmt.Println(i)
    i++
}

// infinite loop
for {

    fmt.Println(i)
    i++
}
```


### exiting early

- We can use `continue` statement to leave loop early and start over

- We can use `break` if you want to exit loop completely on certain condition. Nested loops need special handling to `break` out, and that is `labels`

```go
Loop:
for i :=0 ; i<3 ; i++{
    for j :=0 ; j<3 ; j++{
        if true{
            break Loop // this jumps out of loop
            // OR continue labelName
        }
    }
}

```

### looping through collection

- Looping through collection:
```go

for k, v := range collectionName{

}

```

- For slices:
```go
s := []int{1,2,3}
for k,v := range s{
    fmt.Println(k,v) // prints index and value, in case of maps, prints key:values
}
```

- For **strings**, we get `key` as `index` and `value` as `unicode` value for string character.


- You can skip using either `key` or `value`:

```go
for k := range collectionName{ // just keys

}

// OR 
for _, v := range collectionName{ // just values

}
```


# Defer, Panic and Recover

## Defer
- `defer` will, as the name suggests, defers the execution of a statement to last. It has _LIFO_ order, like a Stack. It is often good to use `defer` to close _connections_ or _files_.

- Caveat:

```go

func main(){
    a := "start"
    defer fmt.Println(a)
    a = "end"
}
// This prints out `start` as the defer uses `eager` evaluation and stores that in memory for the arhuments. So `a= "end"` doesn't affect the deferred print.
```


## Panic

- `panic()` is a function that you can use to `panic` or `raise exception` within your code when you need to raise exceptions.

- Go doesn't always have opinion if something is bad enough for it to `panic` and halt execution, its upto golang developer to decide that and raise panic.

- `Deferred` statements/functions always execute even when a `panic` happens.

```go

func main(){
    a := "start"
    defer fmt.Println(a)
    panic("Bad things happened")
    a = "end"
}

/*  Prints:
start
panic: Bad things happened

*/
```

- Defer takes function calls, not a function:

```go

func main(){
    fmt.Println("start")
    defer func(){ // deferred anonymous function
        // do something 
        fmt.Println("In deferred func")
    }() // don't forget to call it
    panic("something bad happened")
    fmt.Println("end")
}


/* Output:
start
in deferred func
something bad happendd
*/
```

##  Recover


- We can use `revocer()` function to handle `panic` situation.
- When application doesn't `panic`, `recover()` will return `nil` otherwise it returns the error that caused the panic.

```go

func main(){
    fmt.Println("start")
    defer func(){ // deferred anonymous function
	if err:= recover(); err !=nil {
		// do some error handling
		log.Println("Error:", err)
	}
        fmt.Println("In deferred func")
    }() // don't forget to call it
    panic("something bad happened")

}
/* output:
start
2009/11/10 23:00:00 Error: something bad happened
In deferred func
*/

```

- If you had another function outside of `main()` that `panic` and then `recovers`, the function would stop executing but `main()` would still execute.
- If you want to continue `panic` again, then you should `recover()` and then call `panic()` again and if calling function was `main()`, main would panic too.




# Pointers


## Creating & Dereferencing pointers 

- `Pointers` are special variables that point to address of another variable.

- `&` is called address of operator while `*` is called dereferencing operator that dereferences the pointer.

- Dereferencing operator `*` has _lower precendence_ than `.` operator.
```go

func main(){

    var a int = 42
    var b *int = &a // pointer is of type int hence *int and stores &a - which means address of variable a
    fmt.Println(&a, b) // prints address of a and value of b, which is again address of a
    fmt.Println(a,*b) // prints value of a and value of a as *b points to value of a
    a = 27
    fmt.Println(a, *b) // 27 27
    *b = 14
    fmt.Println(a, *b) // 14 14
}
```

- __Pointer Arithmatic__ is __not allowed__ in GoLang by default. If you absolutely need it, look at package `unsafe` in golang.
```go
a := [3]int{1, 2, 3}
b := &a[0]
c := &a[1] - 4 // this causes error as you can't do arithmatic operations on memory addresses, which you may be able to do in `C`. That's because GoLang believes in simplicity and pointer arithmatic can be fairly complex code

```

## The new function

- `new` function can be used to initialize structs. But in this case, it will
  always cause fields to be `zero value` initialized.

```go
func main(){
    var ms *myStruct
    fmt.Println(ms) // prints <nil>
    ms = new(myStruct)
    fmt.Println(ms) // &{0}
    (*ms).foo = 42 // Dereferencing operator `*` has _lower precendence_ than `.` operator. hence use `()` to get higher prcendence
    fmt.Println((*ms).foo) // this can be changed to `ms.foo` and Golang will work it out itself

    // (*ms).foo == ms.foo in this case

}
type myStruct struct{
    foo int
}
```

## Working with nil

- An empty pointer has a special zero value of  `<nil>`

## Types with internal pointers

- Array, struct copy is `value` type, while slices are `reference` or `pointer` type.
- Maps also have a `pointer` to underlying data and hence are `reference` type.


# Functions

- GoLang always has to have `package main` and has to have `func main()` that doesn't return anything.
- Functions are declared with `func` keyword.
- Names are `PascalCase` or `camelCase` for distinguishing private vs public functions.
- Function parameters are added between `()` as `func name(varName type)`
- If function returns values are appended at the end as `func name(varName type) returnType`
- Multiple parameters can be passed as `func name(v1 type, v2 type)`
- If all parameters are of same type `func name(v1, v2 type)` is allowed
- If you change value of variable within a function, the variable will not be modified in calling function, as golang passes values by copy, unless the variable passed is a pointer or reference type.
- Passing a large datastructure is better to be passed by pointers as long as you are careful, so that it gives you performance improvements.
- If a function can take variable number of parameters, you can achieve that using `func name(varName ....type)` telling the compiler that this function takes variable number of args of type `type`, we receive data as slice.
- The variable number of parameters as shown above are called _variadic params_ and they always have to be at the end of function signature, as `func
(var1 type, var2 type, var3 ...type)` where `var3`  is a variadic parameter.
- Golang allows return value to be `*type` and function can return `&data` and calling function can dereference it using `*data`. In many other languages, this maybe a risky operation as called function memory stack would be deleted once execution is finished and you will be pointing to the data that doesn't exists or been deleted, but in the GoLang this kind of return is recognized by go runtime and data is copied into heap - which is shared memory storage accessible to your program.
- _Named return values_ is another feature that you have, where you can do:
```go
func name(var1 type) (result type){
    // you have access to `result` as local variable
    return // no need to write result as name here, go runtime knows you are returning result although this can be confusing for the reading in large/long functions
}
```

- We can have multiple return values as:
```go
// this is very common/idiomatic way to write functions
func divide(a, b float64)(float64, error){
    if b == 0.0 {
        return 0.0, fmt.Errorf("Cannot divide by zero")
    }
    return a/b, nil
}

```

## Anonymous functions

- This kind of function is created once and used once, that's it.
```go
func main(){

    func(){
        fmt.Println("Hello from anonymous function")
    }() // call it immediately, if not passed, compiler says "Function evaluated but not used"
    i := 42
    func(i int){
        fmt.Println("Inside anon func with `i`", i)
    }(i) // we can also have params
    // assign to variable
    f := func (){
        fmt.Println("Hello go ")
    }
    // assign to variable - verbose version
    var f2 func() := func (){
        fmt.Println("Hello go ")
    }

    var divide func(int, int)(int, error)
    divide = func(a,b int) (int, error){
        // do things
    }
    d, err := divide(25,5) // this is how we will call divide function
    if err != nil{
        return
    }
    fmt.Println(d)
}
```


## Methods

- This are functions that are called on an instance/object/variable of some type struct.

```go
func main(){
    g := greeter{
        greeting: "Hello",
        name: "Go"
    }
    g.greet()
}

type greeter struct{
    greeting string
    name string
}

func (g greeter) greet(){ // value receiver as `g` is `greeter` not `*greeter`, so you receive copy of greeter object and not actual object
    fmt.Println(g.greeting, g.name)
}


func (g *greeter) greet(){ // pointer receiver as `g` is `*greeter` so you receive actual object and any modifications made here reflect in original object, although we will still call it as g.greet()
    fmt.Println(g.greeting, g.name)
}
```


# Interfaces 
(Note: Need to learn and practice more of this topic)

- Interfaces are a type, we use `interface` keyword to define it.
- Interface do not represent data, but represent methods.
- Interface name should be suffixed with an `er`. In golang there are many single-method interfaces and those are named after method they implement. So for interface with `Write` method, name would be `Writer` for that interface.
- Implementing interfaces is an implicit operation in Golang. It provides polymorphic behavior.
- Example1:
```go
type Writer interface{
    //Method(input type)(return type)
    Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error){
    n, err := fmt.Println(string(data))
    return n, err
}

func main(){
    var w Writer = ConsoleWriter{}
    w.Write([]byte("Hello Go!"))
}

```

- Example2, you don't have to use struct to implement interfaces, any type that can have methods can be used:
```go

type Incrementer interface {
    Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int{
    *ic++
    return int(*ic)
}

func main(){
    myInt := IntCounter(0)
    var inc Incrementer = &myInt
    for i:=0; i<10; i++{
        fmt.Println(inc.Increment())
    }
}

```

- You cannot directly have interface that is of type `int` or any other primitive type, as that is language builtin type and you can't modify it, you have no control over it.

- `io.Writer` interface is one of the most common interfaces in GoLang.

- Composed Interfaces example:
```go

package main

import "bytes"

type Writer interface {
	//Method(input type)(return type)
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}

type WriterCloser interface{
    Writer
    Closer
}
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bw *BufferedWriterCloser) Write(data []byte) (int, error) {
	// do something
	return 0, nil
}

func (bw *BufferedWriterCloser) Close() error {

	// do something
	return nil
}

//constructor
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
func main() {
	var w WriterCloser = NewBufferedWriterCloser()
    w.Write([]byte("Hello"))
    w.Close()
}

```

- Empty interface is `interface{}` and it has no methods on it. This can be useful, working with multiple things that are not compatible with one another, you need to apply some logic to determine what it exactly is at later point. EmptyInterfaces are useful but you need to type convert those before you can use them in many cases.

-  Type based switch is usually tied with an Empty Interface to be able to decide what type of value is stored in the interface:

```go

func main(){
    var i interface{} =1 // or "1" or true
    switch i.(type){
        case int:
        // do something
        case float64:
        // do something
        case string:
        // do something
        case [3]int: // only matches array with 3 elements of int type
        // do something
        case []int:
        // do something
        default:
        // do something
    }
}

```


- __Need to Study__ : How do interfaces behave when they have valueReceiver vs pointerReceivers in their method sets. Also type conversions.

## Best Practices

- Use many small interfaces over large interfaces, some examples are `io.Writer, io.Reader, interface{}` - these are most common & powerful interfaces in Go Lang. Use composite interfaces.
- Don't export interface for types that will be consumed, it allows user to create and implement their own interfaces for types struct.
- Do export interfaces for types that will be used by package, this is possible because unlike Java/C#, Go lang has implicit implementation of interfaces, you can defer implementation. It allows users to provide their own implementation to the interfaces when they use your library.
- Design functions and methods to receive interfaces whenever possible instead of concrete types.



# GOROUTINES - Concurrency in GoLang

## Create GoRoutines

- `go` keyword before a function call instructs golang to spin off a new "green" thread.

- Most tradtional programming languages used OS threads which means they have individual call stack for each thread and those are very very large.

- In GoLang, it follows something that was seen in Erlang. It is called `Green Thread`. This model would be beneficial as opposed to other languages where creation of thread could be expensive operation, GoLang has a scheduler that takes turns for each GoRoutine and maps it to OS threads for a period of time.

- GOROUTINE is basically abstraction of a thread.
- GoROUTINES start with very small stack spaces, very cheap to create/destroy and helps us not having to deal with low level OS threads.

- `main()` function itself also runs as a GOROUTINE.

```go

package main

func main(){
    go sayHello() // `go` makes it a goroutine
    // This code exits before "Hello" could be printed as
    // after spawning GO routine, application had nothing else to do.

}

func sayHello(){

    fmt.Println("Hello")
}

```
- Another example of GOROUTINE is using it with an anonymous function: 
```go

func main(){

    var msg = "Hello"
    go func (){
        fmt.Println(msg)
    }()
    time.Sleep(100 * time.Millisecond) // this avoids from main function exiting too quickly before giving GOROUTINE a chance to execute, although it is not a best practice
}
```

- Creating dependencies on variables outside goroutines can cause issues:

```go

func main(){

    var msg = "Hello"
    go func (){
        fmt.Println(msg)
    }()
    msg = "goodbye!"
    time.Sleep(100 * time.Millisecond) // this avoids from main function exiting too quickly before giving GOROUTINE a chance to execute, although it is not a best practice
}

// OUTPUT:
// goodbye

// This maybe called a Race Condition, and its bad. This happens because goroutine maybe executed slightly after `msg` is modiefied
// go scheduler may not interrupt main function even if Goroutine was spawned
```

- To fix race condition above, we could do:
```go

func main(){

    var msg = "Hello"
    go func (msg string){ // here we pass by value, a copy of original msg variable
        fmt.Println(msg)
    }(msg)
    msg = "goodbye!" // hence this doesn't override msg in the anonymous function 
    time.Sleep(100 * time.Millisecond) // this avoids from main function exiting too quickly before giving GOROUTINE a chance to execute, although it is not a best practice
}
// OUTPUT:
// hello
```


## Synchronizations


- To avoid the problem of application exiting before GOROUTINE is executed, we could use following instead of `time.Sleep()`: 

```go

var wg = sync.WaitGroup{} // from sync package

func main(){
    var msg = "Hello"
    wg.Add(1) // adding 1 to waitGroup for 1 goRoutine we are about to spawn
    go func(msg string){
        fmt.Println(msg)
        wg.Done() // this will decrement value in wait group, in this case 1 - 1 = 0
    }
    msg = "goodbye"
    wg.Wait() // this will wait for all the Goroutines to decrement waitgroup object until its down to 0

}
```

-  Another little more complex example of Goroutines waitGroups sync is:
```go

var wg = sync.WaitGroup{}
var counter = 0

func main(){
    for i:=0; i<10; i++{
        wg.Add(2)
        go sayHello()
        go increment()
    }
    wg.Wait()
}


func sayHello(){
    fmt.Printf("Hello $%v\n", counter)
    wg.Done()
}

func increment(){
    counter++
    wg.Done()
}

/*Output:

Hello $1
Hello $2
Hello $3
Hello $8
Hello $5
Hello $5
Hello $7
Hello $10
Hello $9

// Order of printing above is not reliable as both functions are spawned as goroutines and they keep running against one another in a race condition.
*/
```

- Synchronizing with Mutex(Mutually Exclusive locks):
```go

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // Allows multiple readers, but only one writer
                       // if writer wants to write, it has to wait for all readers to be done reading.

func main(){
    for i:=0; i<10; i++{
        wg.Add(2)
        go sayHello()
        go increment()
    }
    wg.Wait()
}


func sayHello(){
    m.RLock() // ReadLock
    fmt.Printf("Hello $%v\n", counter)
    m.RUnlock() // ReadUnlock
    wg.Done()
}

func increment(){
    m.Lock() // Write Lock
    counter++
    m.Unlock() //Write Unlock
    wg.Done()
}

/* Output:
//Run1 could get
Hello $1
Hello $2
.
.
.
.Hello $2

// Run2 could get 


Hello $1
Hello $2
Hello $3
Hello $3
Hello $4
Hello $4
Hello $4
.
.
.

// As you can see order is preserved but output is still unpredictable. This is because Mutexes are still within Go Routine and Go Routines don't have particular order of execution.
*/ 

```

- To fix this we can modify the Mutexes and moved out of go routines(although this way we are kind of destroying concurrency by forcing synchronization, this may perform worse than not having any goroutines at all):

```go

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // Allows multiple readers, but only one writer
                       // if writer wants to write, it has to wait for all readers to be done reading.

func main(){
    for i:=0; i<10; i++{
        wg.Add(2)
        m.RLock() // ReadLock
        go sayHello()
        m.Lock() // Write Lock
        go increment()
    }
    wg.Wait()
}


func sayHello(){
    fmt.Printf("Hello $%v\n", counter)
    m.RUnlock() // ReadUnlock
    wg.Done()
}

func increment(){
    counter++
    m.Unlock() //Write Unlock
    wg.Done()
}

/* OUTPUT:

Hello $0
Hello $1
Hello $2
Hello $3
Hello $4
Hello $5
Hello $6
Hello $7
Hello $8
Hello $9

// Now we can re-run this over and over again and order is reliably preserved
*/
```

## Parallelism

- `GOMAXPROCS` is a tunable in Golang that you can use. Minimum it can be set to 1 thread per core. But if you need to go high you could. But going too high can cause other problems, like you may not have enough memory to spawn 10s or 100s of threads.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // this sets maxprocs to only 1 thread and runs application single threaded way
	fmt.Printf("Threads: %v \n", runtime.GOMAXPROCS(-1)) // prints 1, but if we remove previous line, we get output equal number of cores on the system, passing "-1" as arg is just letting us investigate the current value set for GOMAXPROCS
}

```




## Best Practices

- GOROUTINES in GoLang are very powerful, but don't create goroutines in libraries, let consumer control concurrency.
- When you create a goroutine, know how it will end, avoid suble memory leaks. If you have a goroutine that doesn't end will keep eating resources forever.
- Check for race conditions at compile time. use `go run/build` commands with `-race` flag.