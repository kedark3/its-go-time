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
	var clock [8]placeholder
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
		printClock(clock[:])
		fmt.Println()
		clock = [8]placeholder{}
		time.Sleep(time.Second)
	}

	// var zero = placeholder{
	// 	"███",
	// 	"█ █",
	// 	"█ █",
	// 	"█ █",
	// 	"███",
	// }

	// var one = placeholder{
	// 	"██ ",
	// 	" █ ",
	// 	" █ ",
	// 	" █ ",
	// 	"███",
	// }

	// var two = placeholder{
	// 	"███",
	// 	"  █",
	// 	"███",
	// 	"█  ",
	// 	"███",
	// }

	// var three = placeholder{
	// 	"███",
	// 	"  █",
	// 	"███",
	// 	"  █",
	// 	"███",
	// }

	// var four = placeholder{
	// 	"█ █",
	// 	"█ █",
	// 	"███",
	// 	"  █",
	// 	"  █",
	// }

	// var five = placeholder{
	// 	"███",
	// 	"█  ",
	// 	"███",
	// 	"  █",
	// 	"███",
	// }

	// var six = placeholder{
	// 	"███",
	// 	"█  ",
	// 	"███",
	// 	"█ █",
	// 	"███",
	// }

	// var seven = placeholder{
	// 	"███",
	// 	"  █",
	// 	"  █",
	// 	"  █",
	// 	"  █",
	// }

	// var eight = placeholder{
	// 	"███",
	// 	"█ █",
	// 	"███",
	// 	"█ █",
	// 	"███",
	// }

	// var nine = placeholder{
	// 	"███",
	// 	"█ █",
	// 	"███",
	// 	"  █",
	// 	"███",
	// }

	// var colon = placeholder{
	// 	"   ",
	// 	" ░ ",
	// 	"   ",
	// 	" ░ ",
	// 	"   ",
	// }
	// digits := [...]placeholder{
	// 	zero, one, two, three, four, five, six, seven, eight, nine,
	// }
	// for shift := 0; ; shift++ {
	// 	// we need to clear the screen here.
	// 	// or the previous character will be left on the screen
	// 	//
	// 	// alternative: you can fill the rest of the missing placeholders
	// 	//              with empty lines
	// 	fmt.Print("\033[H\033[2J") // clears screen before printing next iteration

	// 	now := time.Now()
	// 	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	// 	clock := [...]placeholder{
	// 		digits[hour/10], digits[hour%10],
	// 		colon,
	// 		digits[min/10], digits[min%10],
	// 		colon,
	// 		digits[sec/10], digits[sec%10],
	// 	}

	// 	for line := range clock[0] {
	// 		l := len(clock)

	// 		// this sets the beginning and the ending placeholder positions (indexes).
	// 		// shift%l prevents the indexing error.
	// 		s, e := shift%l, l

	// 		// to slide the placeholders from the right part of the screen.
	// 		//
	// 		// here, we assume that as if the clock's length is double of its length.
	// 		// this makes things easy to manage: that's why: l*2 is there.
	// 		//
	// 		// shift is always increasing, for it's to go beyond the clock's length,
	// 		// it should be equal or greater than l*2, right (after the remainder of course)?
	// 		//
	// 		// so, if the clock goes beyond its length; this code detects that,
	// 		// and resets the starting position to the first placeholder's index,
	// 		// and it keeps doing so until the clock is fully displayed again.
	// 		if shift%(l*2) >= l {
	// 			s, e = 0, s
	// 		}

	// 		// print empty lines for the missing place holders.
	// 		// this creates the effect of moving placeholders from right to left.
	// 		//
	// 		// l-e can only be non-zero when the above if statement runs.
	// 		// otherwise, l-e is always zero, because l == e.
	// 		//
	// 		// this is one of the other benefits of assuming the length of the
	// 		// clock as the double of its length. otherwise, l-e would always be 0.
	// 		for j := 0; j < l-e; j++ {
	// 			fmt.Print("     ")
	// 		}

	// 		// draw the digits starting from 's' to 'e'
	// 		for i := s; i < e; i++ {
	// 			next := clock[i][line]
	// 			if clock[i] == colon && sec%2 == 0 {
	// 				next = "   "
	// 			}

	// 			fmt.Print(next, "  ")
	// 		}
	// 		fmt.Println()
	// 	}

	// 	time.Sleep(time.Second)
	// }

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
