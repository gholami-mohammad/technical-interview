package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"ti/datasource"
	"time"
)

func ReadCSV(filename string, ch chan datasource.User, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i, d := range data {
		if i == 0 {
			continue
		}
		var u datasource.User
		u.FirstName = d[0]
		u.LastName = d[1]
		u.Email = d[2]
		dateString := d[3]
		date, err := time.Parse("2006/01/02", dateString)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		u.Birthdate = date
		ch <- u

	}
	close(ch)
}

func CountUsersAbove30(ch chan datasource.User, wg *sync.WaitGroup) int {
	defer wg.Done()
	usersAbove30 := 0
	for user := range ch {
		age := time.Now().Year() - user.Birthdate.Year()
		if time.Now().YearDay() < user.Birthdate.YearDay() {
			age--
		}
		if age > 30 {
			usersAbove30++
		}
	}
	return usersAbove30
}

func main() {
	var wg sync.WaitGroup
	userChan := make(chan datasource.User)
	wg.Add(1)
	go ReadCSV("storage/user.csv", userChan, &wg)

	go func() {
		wg.Wait()
		close(userChan)
	}()

	for user := range userChan {
		if time.Now().Year()-user.Birthdate.Year() > 30 {
			fmt.Println("User above 30:", user.FirstName, user.LastName)
		}
	}

	go func() {
		usersAbove30 := CountUsersAbove30(userChan, &wg)
		fmt.Println("users above 30: ", usersAbove30)
	}()

}
