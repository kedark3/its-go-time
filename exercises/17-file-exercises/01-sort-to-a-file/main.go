package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	sort.Strings(args)

	var size int // decide size for allocating the byte array once, avoid reallocation on every append
	for _, s := range args {
		size += len(s)
	}

	data := make([]byte, 0, size+len(args))

	for _, s := range args {
		data = append(data, s...) // s... because a string is slice of bytes, we need to append that to byte slice byte by byte
		data = append(data, '\n')
		fmt.Printf("address of slice %p \n", data)
	}
	err := ioutil.WriteFile("sorted.txt", data, 0777)
	if err != nil {
		fmt.Println("Some error occurred", err.Error())
	}
}
