package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string,error) {
    // if there is no name, then show the error

    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
    formats := []string{
        "hi, %v",
        "whaddub %v",
        "yo yo %v",
    }
    return formats[rand.Intn(len(formats))]
}