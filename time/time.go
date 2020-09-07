package main

/* This import format is called "Factored" import, you could write them on separate lines too as
import "math"
import "fmt"

*/
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welome to the playground!")
	// Fun Fact: This returns 2009-11-10 in the Go Playground!!
	fmt.Println("The time is", time.Now())

	t := time.Date(2020, time.September, 6, 22, 6, 0, 0, time.UTC)
	fmt.Printf("This program was written on: %s \n", t)

	fmt.Println("Tomorrow's date is :", t.AddDate(0, 0, 1))

	// create new time format
	// use format 1,2,3,4,5,6,7 - i.e. Monday, January(1), 2nd(2), 3(3):04(4) PM , 5sec(5), 2006year(6), Mountain-TimeZone(7)
	longFormat := "Monday, January 2, '06 : EST"
	fmt.Println("Tomorrow is", t.AddDate(0, 0, 1).Format(longFormat))
}
