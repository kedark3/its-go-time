package main

import (
	"fmt"
	"time"
)

func main() {
	// This loop prints all digits vertically
	// for _, d := range digits {
	// 	for _, line := range d {
	// 		fmt.Println(line)
	// 	}
	// 	fmt.Println()
	// }
	// This loop prints all digits horizontally
	// ███  ██   ███  ███  █ █  ███  ███  ███  ███  ███
	// █ █   █     █    █  █ █  █    █      █  █ █  █ █   █
	// █ █   █   ███  ███  ███  ███  ███    █  ███  ███
	// █ █   █   █      █    █    █  █ █    █  █ █    █   █
	// ███  ███  ███  ███    █  ███  ███    █  ███  ███
	var clock [10]placeholder
	for {
		// 	\033 is a special control code:
		// [2J clears the screen and the cursor
		// [H moves the cursor to 0, 0 screen position
		fmt.Print("\033[H\033[2J") // clears screen before printing next iteration
		hr, min, sec := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		if sec%10 == 0 {
			fmt.Println("ALARM")
		}
		clock[0], clock[1] = digits[hr/10], digits[hr%10]
		clock[3], clock[4] = digits[min/10], digits[min%10]
		clock[6], clock[7] = digits[sec/10], digits[sec%10]
		clock[8], clock[9] = dot, digits[time.Now().Nanosecond()/1e8]
		if sec%2 != 0 {
			clock[2], clock[5] = separator, separator
		} else {
			clock[2], clock[5] = blankSeparator, blankSeparator
		}
		printClock(clock[:])
		fmt.Println()
		time.Sleep(time.Second / 10)
	}

}

func printClock(digits []placeholder) {

	for line := range digits[0] {
		// Print a line for each placeholder in digits
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}
}
