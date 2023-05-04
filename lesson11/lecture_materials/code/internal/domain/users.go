package domain

import "time"

type User struct {
	ID       int
	Name     string
	Birthday time.Time
}
