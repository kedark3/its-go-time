package main

import (
	"fmt"
	"unsafe"
)

func main() {

	ages := []int{25, 26, 27}

	fmt.Println(cap(ages), len(ages), ages)

	ages1 := ages[:1]
	fmt.Println(cap(ages1), len(ages1), ages1) // cap: 3, len: 1, items: [25]

	ages13 := ages[1:3]
	fmt.Println(cap(ages13), len(ages13), ages13) // 2 2 [26 27]

	fmt.Println(ages13[0]) // sees 26 as 0th index element, not 25, its not visible to ages13 slice

	fmt.Println(ages1[0]) // yields 25
	// fmt.Println(ages1[1]) // will not work, even though cap is 3, its len is 1 so it can't access anything beyond 0th index
	// fmt.Println(ages1[2]) // will not work, even though cap is 3, its len is 1 so it can't access anything beyond 0th index

	ages13[1] = 189               // modifies the backing array for ages slice and change is reflected in ages, ages13 slices
	fmt.Println("ages", ages)     // changed
	fmt.Println("ages1", ages1)   // unchanged
	fmt.Println("ages13", ages13) // changed
	S := []string{}               // Is initialized so it’s not nil
	var s []string                // Is a nil slice and it is uninitialized

	if S != nil { // you can compare a slice with nil
		fmt.Println("Capital S slice is not nil")
	}
	if s == nil {
		fmt.Println("lower case s slice is nil")
	}

	// if s == S {}  // this does not work as you can't compare slices to one another
	if len(s) == len(S) { // we can compare slices this way - to check if they are equal in size
		fmt.Println("Slices are both equal")
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // creates a copy for arr1 for arr2
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr2[1] = 324     // doesn't change arr1
	fmt.Println(arr1) // [1 2 3]
	fmt.Println(arr2) // [1 324 3]

	s1 := []int{1, 2, 3}
	s2 := s1 // creates a copy for s1 for s2
	fmt.Println(s1)
	fmt.Println(s2)
	s2[1] = 324     // changes s1 as well
	fmt.Println(s1) // [1 324 3]
	fmt.Println(s2) // [1 324 3]

	data := []string{"I", "am", "KK"}
	fmt.Println(data)
	change(data)
	fmt.Println(data)

	fmt.Println("Size of Slice data", unsafe.Sizeof(data)) // all slices contain 3 fields in the slice header,
	fmt.Println("Size of Slice s1", unsafe.Sizeof(s1))     // len, capacity, and ptr to backing array, 8 bytes each
	// hence copying slices and passing it to the functions is a cheap operation
	// size of a slice is always 24 bytes

	var array [10]string
	fmt.Println(array[:5], len(array[:5]), cap((array[:5]))) // `array[:5]` returns a slice with the first 5 elements of the `array` (len is 5), but there are 5 more elements in the backing array of that slice, so in total its capacity is 10.

	// var tasks []string  // this is have len:0, pointer:0, cap:0

	//`array` is 1000 x int64 (8 bytes) = 8000 bytes. Assigning an array copies all its elements,
	// so `array2` adds additional 8000 bytes. A slice doesn't store anything on its own.
	// Here, it's being created from array2, so it doesn't allocate a backing array as well.
	// A slice header's size is 24 bytes. So in total: This program allocates 16024 bytes.
	var arrayFloat [1000]int64

	array2 := arrayFloat
	slice := array2[:]
	fmt.Println(unsafe.Sizeof(array2), unsafe.Sizeof(slice))

}

func change(data []string) { // go creates copy of original slice and sends it to the function
	// but it does not copy the backing array, it still shares the same backing array between calling function
	// and this function
	data[2] = "newData added by change function call" // this changes the original slice backing array
}

/* Output1:
3 3 [25 26 27]
3 1 [25]
2 2 [26 27]
26
25
panic: runtime error: index out of range [1] with length 1

goroutine 1 [running]:
main.main()
        /home/kkulkarni/go/src/github.com/kedark3/its-go-time/exercises/16-slices/00-slice-playground/main.go:20 +0x312
exit status 2
*/

/*Output2:
3 3 [25 26 27]
3 1 [25]
2 2 [26 27]
26
25
ages [25 26 189]
ages1 [25]
ages13 [26 189]
Capital S slice is not nil
lower case s slice is nil
Slices are both equal
*/
