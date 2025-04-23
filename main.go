package main

import (
	"fmt"
)

func main() {
	// define user channel
	// userChan := ?
	usersAbove30 := 0

	go func() {
		// TODO: 1- read from multiple data sources
		// read csv file and log file line by line at same time
		// send parsed user data to userChan
	}()

	go func() {
		// TODO: 2- read users on userChan
		// calculate age
		// count if older than 30
	}()

	// TODO: 3- fix deadlock issue if happened

	fmt.Println("users above 30: ", usersAbove30)
}
