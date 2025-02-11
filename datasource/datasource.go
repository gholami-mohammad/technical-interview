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

type DataFeed interface {
	ReadUsers(chan User) error
}

type DataSource struct {
	dataFeeds []DataFeed
}

// TODO: create a function called "Stream" and read users from all defined data feeds
