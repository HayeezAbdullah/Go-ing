package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string, fathername string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v, son of %v", name, fathername)
    return message
}