package greetings

import "fmt"

// Hello returns a greeting for the named person.

func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}


/*
< := > 
operator creates the value of the right side type (like an r-value) (initualizes)
like = then assigns the value.

alternative:
	string message
	message =
*/
