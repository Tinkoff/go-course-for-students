package main

import (
	"fmt"
	"time"
)

type User struct {
	ID int
}

type UserService interface {
	FindUser(id int) *User
}

type userService struct{}

func (us *userService) FindUser(id int) *User {
	return &User{ID: id}
}

func main() {
	var us UserService = &userService{}

	us = &loggedUserService{us}

	us = &timeitUserService{us}

	fmt.Printf("User: %+v\n", us.FindUser(42))
}

type loggedUserService struct {
	us UserService
}

func (lus *loggedUserService) FindUser(id int) *User {
	fmt.Println("trying to get user with id =", id)
	return lus.us.FindUser(id)
}

type timeitUserService struct {
	us UserService
}

func (tus *timeitUserService) FindUser(id int) *User {
	started := time.Now()
	defer func() {
		fmt.Println("time:", time.Now().Sub(started))
	}()
	return tus.us.FindUser(id)
}
