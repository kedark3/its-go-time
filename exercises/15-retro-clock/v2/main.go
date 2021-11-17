package main

import (
	"fmt"
	"time"
)

type placeholder [5]string

func main() {
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	separator := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}
	blankSeparator := placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}
	digits := [11]placeholder{zero, one, two, three, four, five, six, seven, eight, nine, separator}
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
	var clock [16]placeholder
	for i := 8; i < 16; i++ {
		clock[i] = blankSeparator
	}
	i, j, switchSide := 0, 0, false
	for {
		// 	\033 is a special control code:
		// [2J clears the screen and the cursor
		// [H moves the cursor to 0, 0 screen position
		fmt.Print("\033[H\033[2J") // clears screen before printing next iteration
		hr, min, sec := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		clock[0], clock[1] = digits[hr/10], digits[hr%10]
		clock[3], clock[4] = digits[min/10], digits[min%10]
		clock[6], clock[7] = digits[sec/10], digits[sec%10]
		if sec%2 != 0 {
			clock[2], clock[5] = separator, separator
		} else {
			clock[2], clock[5] = blankSeparator, blankSeparator
		}
		if i == 8 {
			switchSide = true
			j = i
		} else if i == 0 {
			switchSide = false
		}
		if !switchSide {
			printClock(clock[i:8])
			i++
		} else {
			printClock(append(clock[j:15], clock[:j-8]...))
			i--
			j++
		}
		fmt.Println()
		time.Sleep(time.Second)
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
