package user

import "fmt"

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (u *User) String() string {
	return fmt.Sprintf("<User(id=%d, name=`%s`)>", u.ID, u.Name)
}

func New(id int64, name string) User {
	return User{ID: id, Name: name}
}
