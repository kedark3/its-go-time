package main

import (
	"fmt"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {

	houses := map[string][]string{
		"gryffindor": {
			"weasley",
			"hagrid",
			"dumbledore",
			"lupin",
		},
		"hufflepuf": {
			"wenlock",
			"scamander",
			"helga",
			"diggory",
		},
		"ravenclaw": {
			"flitwick",
			"bagnold",
			"wildsmith",
			"montmorency",
		},
		"slytherin": {
			"horace",
			"nigellus",
			"higgs",
			"scorpius",
		},
		"bobo": {
			"wizardry",
			"unwanted",
		},
	}

	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}
	students, ok := houses[args[0]]
	if !ok {
		fmt.Printf("Sorry. I don't know anything about %q.\n", args[0])
		return
	}
	// Following ops can modify original student slice within the map as
	// students variable is of type slice and it points to same backing array as the students slice in original map
	// students[0] = "KK"

	// Cloning can be done as :
	clone := make([]string, len(students)) // this causes zero-ing of elements upto the len
	copy(clone, students)
	// OR as :
	// clone := append([]string(nil), students...) // here we create nil slice, so no backing-array no-zero-ing, append creates new backing array

	sort.Strings(clone)
	fmt.Printf("\t ~~~ %s Students ~~~\n", args[0])
	for _, s := range clone {
		fmt.Printf("\t\t + %s \n", s)
	}

	// using Maps as sets
	// Any key that is not in the map would return false by default so we can use this fact to check membership of a key in the set
	// m["kk"] would be true, m["john"] would yield false. So its similar to `"kk" in set` of python
	m := map[string]bool{"kk": true, "tesla": true}
	if m["kk"] && m["tesla"] {
		fmt.Println("KK has Tesla")
	}
}

/* see benchmark results for the same(MakeCopy vs Append cloning methods - both didn't come out as conclusive winner):

kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               635.2 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               651.1 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.132s
 kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               599.7 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               624.2 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.125s
 kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               623.9 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               632.6 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.128s
 kkulkarni@kkulkarni  exercises/16-slices/27-benchmark-slices~   master  go test -bench=. -benchmem -benchtime=100000x
goos: linux
goarch: amd64
pkg: main/exercises/16-slices/27-benchmark-slices
cpu: Intel(R) Core(TM) i7-8665U CPU @ 1.90GHz
Benchmark_MakeCopy_128-8          100000               562.7 ns/op          2048 B/op          1 allocs/op
Benchmark_Append_128-8            100000               630.0 ns/op          2048 B/op          1 allocs/op
PASS
ok      main/exercises/16-slices/27-benchmark-slices    0.122s

*/
