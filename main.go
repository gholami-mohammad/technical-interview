package main

import (
	"fmt"
)

func main() {
	// define user channel
	// userChan := ?
	usersAbove30 := 0

	// TODO: count users older than 30 years old
	// TODO: fix deadlock issue

	go func() {
		// TODO: read from multiple data feeds
	}()

	go func() {
		// TODO: read users on userChan
		// calculate age
		// count if older than 30
	}()

	fmt.Println("users above 30: ", usersAbove30)
}
