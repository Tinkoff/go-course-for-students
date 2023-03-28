package main

// db:table=users
type User struct {
	Name       string `db:"name"`
	Age        string `db:"age"`
	Occupation string `db:"occupation"`
	Options    []string
}
