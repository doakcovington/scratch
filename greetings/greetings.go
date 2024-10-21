package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// returns set of greeting messages which is randomized 
func randomFormat() string  {
    // Array of message formats
    formats := []string{
        "Hi %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[rand.Intn(len(formats))]
}