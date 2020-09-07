package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
)

// Tour struct to extract tour data from JSON in Explore California
type Tour struct {
	Name, Price string
}

func main() {
	content := "Hello World from Go! - Written by Go in this file"

	file, err := os.Create("./goFile.txt")
	checkError(err)

	defer file.Close()
	ln, err := io.WriteString(file, content)

	checkError(err)

	fmt.Println("Written", ln, "Chars to the file")

	readFromFile, err := ioutil.ReadFile("./goFile.txt")
	// without type casting in string() the readFromFile will print raw bytes.
	fmt.Println(string(readFromFile))

	// Read from an URL
	url := "https://www.redhat.com/en"

	response, err := http.Get(url)
	checkError(err)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(bytes[:100]))

	// Read JSON data from a Server
	url = "http://services.explorecalifornia.org/json/tours.php"
	content = contentFromServer(url)
	tours := tourFromJSON(content)
	fmt.Println(tours)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToNearestEven)
		fmt.Printf("\n %v ($%.2f)", tour.Name, price)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	response, err := http.Get(url)
	checkError(err)
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)

	return string(bytes)
}

func tourFromJSON(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	checkError(err)

	var tour Tour

	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours

}
