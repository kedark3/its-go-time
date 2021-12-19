package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + Get all the files in a directory using ioutil.ReadDir
//     (A directory is also a file)
//
//   + You can use IsDir method of a FileInfo value to detect
//     whether a file is a directory or not.
//
//     Check out its documentation:
//
//     go doc os.FileInfo.IsDir
//
//     # or using godocc
//     godocc os.FileInfo.IsDir
//
//   + You can use '\t' escape sequence for indenting the subdirs.
//
//   + You can find a sample directory structure under:
//     solution/ directory
//
// ---------------------------------------------------------

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}
	var subdirs []byte
	for _, dir := range items {
		list, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}
		subdirs = append(subdirs, []byte(fmt.Sprintf("%s\n", dir))...)
		for _, f := range list {
			if f.IsDir() {
				subdirs = append(subdirs, []byte(fmt.Sprintf("\t%s/\n", f.Name()))...)
			}
		}
	}
	err := ioutil.WriteFile("subdirs.txt", subdirs, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
