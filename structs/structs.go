// Time for some  C language nostalgia again with Structs

package main

import "fmt"

type car struct {
	Make  string
	Model string
	Year  string
	Price int
}

type raceCar struct {
	typeOfRace string
	carName    string
}

type vehicle interface {
	Transmission()
}

func (c car) Transmission() {
	fmt.Println("I am a car, so I will always have Auto Transmission. Welcome to sad 2020!")
	fmt.Println("Bye bye manual transmission cars :(")
}

func (rc raceCar) Transmission() {
	fmt.Println("I am a RaceCar, so I can have Manual or Auto Transmission")
}

// The object c is passed to IsExpensive by Value. But if you annotate it with * as *car
// The changes made to car object in this method will modify the original object
func (c car) IsExpensive() {
	if c.Price > 30000 {
		fmt.Println("You will buy it one day :) ", c.Make, c.Model, c.Price)
	} else {
		fmt.Println("That's a great deal!", c.Make, c.Model, c.Price)
	}
}

// This function is using call by reference on car object
// Every time you call the function, Price of car Appreciates(in my fairy land world)
func (c *car) Appreciate() {
	c.Price += c.Price / 10
}

func main() {
	// defer keyword works such that they will be executed at the end of given function
	// every defer statement is put into Stack so they are executed in LIFO order
	defer fmt.Println("I am 1st stattement in main() but I was deferred")
	x := 1000
	defer fmt.Println("Value of X right after initialization(but printing was deferred)", x)
	x += 10
	// defer is useful for closing files or DB connections
	fmt.Println("Value of X right after initialization", x)
	city := vehicle(car{"Honda", "City", "2012", 12000})
	tesla := vehicle(car{"Tesla", "ModelX", "2020", 99000})
	fmt.Println(city)
	// now print with field names
	fmt.Printf("\n%+v\n", city)
	fmt.Printf("\n%+v\n", tesla)
	// city.IsExpensive()
	// tesla.IsExpensive()

	// city.Appreciate()
	// fmt.Printf("\nCar value after appreciation: %+v\n", city)
	// city.Appreciate()
	// fmt.Printf("\nCar value after appreciation: %+v\n", city)

	// tesla.Appreciate()
	// fmt.Printf("\nCar value after appreciation: %+v\n", tesla)
	// tesla.Appreciate()
	// fmt.Printf("\nCar value after appreciation: %+v\n", tesla)

	f1Car := vehicle(raceCar{"Formula1", "Mazda"})
	nfsCar := vehicle(raceCar{"Need For Speed", "McLaren F1"})

	f1Car.Transmission()
	nfsCar.Transmission()
	city.Transmission()
	tesla.Transmission()
}
