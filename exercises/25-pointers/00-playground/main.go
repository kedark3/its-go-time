package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("***** ARRAYS")
	arrays()

	fmt.Println("***** SLICES")
	slices()

	fmt.Println("***** MAPS")
	maps()

	fmt.Println("***** STRUCTS")
	structs()
}

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	// still prints [1,2,3] as
	// incr function can't change the array
	fmt.Printf("array nums : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Printf("array nums by ptr : %p\n", &nums)
	fmt.Println(nums)

}

func incr(nums [3]int) {
	for i := range nums {
		nums[i]++
	}
	// address here for nums shows different memory location
	// than one from arrays(), as its pass by value
	fmt.Printf("incr nums : %p\n", &nums)
	fmt.Println("***** INCR\n", nums)
	// go allows us to get address of slice/array elements individually
	fmt.Printf("up &nums[1] : %p\n", &nums[1])
	fmt.Printf("up &nums[2] : %p\n", &nums[2])
}

func incrByPtr(nums *[3]int) {
	for i := range nums {
		// go internally changes following call to
		// (*nums)[i]++
		nums[i]++
	}
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}
	// slice is passed by value but the slice value is a slice header and
	// it contains pointer to the backing array
	up(dirs)
	fmt.Printf("slices list %p %q \n", &dirs, dirs)
	// passing it this way using pointer is not common, so one should not use it
	// but this is an exmple and you can see doing this will pass entire slice header
	// by reference and called function can modify the slice header and changes are reflected
	// back in the calling function
	upPtr(&dirs)
	fmt.Printf("slices list %p %q \n", &dirs, dirs)
}
func up(list []string) {
	// hence receiving the slice and modifying/updating elements of the slice cause changes
	// to be visible in calling/called functions both
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
	// modifying the slice here doesn't show up in calling function as the backing array is now changed
	list = append(list, "MISSING BUG")
	fmt.Printf("up list %p %q \n", &list, list)
	// go allows us to get address of slice/array elements individually
	fmt.Printf("up &list[1] : %p\n", &list[1])
	fmt.Printf("up &list[2] : %p\n", &list[2])

}

func upPtr(list *[]string) {
	lv := *list
	// hence receiving the slice and modifying/updating elements of the slice cause changes
	// to be visible in calling/called functions both
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}
	// modifying the slice here doesn't show up in calling function as the backing array is now changed
	*list = append(*list, "MISSING BUG")
	fmt.Printf("up ptr list %p %q \n", list, list)
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fmt.Println(confused)
	fix(confused)
	// you can see this will print the correct map
	// because map is also passed by value and value is just the header
	// and it points to backend map storage
	// so if called function modifies the map, it reflects back in here
	fmt.Println(confused)
}
func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}
	fmt.Println(myHouse)
	addRooms(myHouse)
	// Rooms count doesn't change as the struct var is passed by value
	// changes in the called function do not reflect back up in the calling function
	fmt.Println(myHouse)
	// if we call the function again and make it return the new value we can get correct value
	myHouse = addRooms(myHouse)
	fmt.Println(myHouse)

	// but we can do away without having to "return and re-assign" value to
	// override old variable, just passing address of(or pointer to) it and
	// called function can modify values directly stored in that variable
	addRoomsPtr(&myHouse)
	fmt.Printf("%p %v\n", &myHouse, myHouse)

}

func addRooms(h house) house {
	h.rooms += 1
	return h
}

func addRoomsPtr(h *house) {
	// h is a pointer variable of type house that contains the address of the variable
	// it received. So to access the value we need to do *h, but go does that automatically
	(*h).rooms += 1
	h.rooms += 1 // same as: (*h).rooms++ , go does that automatically
	fmt.Printf("%p %v %v\n", h, h, *h)
	// it prints:
	// 0xc00000c0c0 &{My House 8} {My House 8}
	// as you can see above, 2nd h prints & and *h doesn't as it is way of go telling us
	// value of the variable is address of myHouse while in *h we are explicitly retrieving its value

	// go also allows us to access address of the each individual struct field as below
	fmt.Printf("&h.name : %p\n", &h.name)
	fmt.Printf("&h.rooms : %p\n", &h.rooms)
}
