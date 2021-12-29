package main

import (
	"fmt"
	"os"
)

/*
Spam Masker Challenge Tips

👉You can also find the (maybe) updated version of this document in the code repository, here.

Rules:

    You shouldn't use a standard library function.

    You should only solve the challenge by manipulating the bytes directly.

    Manipulate the bytes of a string using indexing, slicing, appending etc.

    Be efficient: Do not use string concat (+ operator).

        Instead, create a new byte slice as a buffer from the given string argument.

        Then, manipulate it during your program.

        And, for once, print that buffer.

    Mask only links starting with http://

    Don't check for uppercase/lowercase letters

        The goal is learning how to manipulate bytes in strings, it's not about creating the perfect masker.

        For example: A spammer can prevent the masker like this (for now this is OK):

            "Here's my spammy page: hTTp://youth-elixir.com"
                                     ^^

    But, you should catch this:

        INPUT:
        Here's my spammy page: http://hehefouls.netHAHAHA see you.

        OUTPUT:
        Here's my spammy page: http://******************* see you.

Steps:

    Check whether there's a command line argument or not. If not, quit from the program with a message.

    Create a byte buffer as big as the argument.

    Loop and detect the http:// patterns

    Copy the input character by character to the buffer

    If you detect http:// pattern, copy the http:// first, then copy the *s instead of the original link until you see whitespace character.

    For example:

        INPUT:
        Here: http://www.mylink.com Click!

        OUTPUT:
        Here: http://************** Click!

    Print the buffer as a string
*/

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me some strings.")
		return
	}
	text := args[0]
	var (
		found bool
		mask  = '*'
		http  = "http://"
		n     = len(http) // len of http
	)

	buf := []byte(text) // don't use string buffer and don't use `+` as every time you use `+` go creates new string value as strings are immutable

	for i := 0; i < len(text); i++ {

		if len(text[i:]) >= n && text[i:i+n] == http {
			i += n
			found = true
		}

		switch text[i] {
		case '\n', '\t', ' ':
			found = false
		default:
			if found {
				buf[i] = byte(mask)
			}
		}
	}
	fmt.Printf("%s", buf)

}
