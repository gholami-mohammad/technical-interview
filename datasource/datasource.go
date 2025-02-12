package datasource

import (
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Birthdate time.Time
}
