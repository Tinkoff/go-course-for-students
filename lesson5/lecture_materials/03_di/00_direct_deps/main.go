package main

import (
	"sync"
)

type UserService struct {
	store *sync.Map
}

func NewUserService() *UserService {
	store := &sync.Map{}

	return &UserService{
		store: store,
	}
}

func main() {
	_ = NewUserService()
}
