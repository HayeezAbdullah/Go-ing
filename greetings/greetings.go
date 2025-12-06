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

// hellos for many people

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message
    }

    return messages, nil
}

func randomFormat() string {
    formats := []string{
        "hi, %v",
        "whaddub %v",
        "yo yo %v",
    }
    return formats[rand.Intn(len(formats))]
}