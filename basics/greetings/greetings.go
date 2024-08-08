package greetings

import (
	"errors"
	"fmt"
)

// Returns a "hello" for a given person
func Hello(name string) (string, error) {
	// If no name was given return an error
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
