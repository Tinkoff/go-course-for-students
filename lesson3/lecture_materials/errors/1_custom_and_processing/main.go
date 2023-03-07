package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ------------------------------------------

type CustomError struct {
	Info string
	Code int
}

func (s CustomError) Error() string {
	return "info: " + s.Info + ", status code: " + strconv.Itoa(s.Code)
}

var errNotFound = CustomError{Code: 404, Info: "user not found"}

// ------------------------------------------

var errAuth = errors.New("can not auth")

// ------------------------------------------

type User struct {
	ID       int
	Name     string
	Password string
}

var localStorage = map[string]*User{
	"t.e.razumov": {
		ID:       0,
		Name:     "t.e.razumov",
		Password: "qweqwe",
	},
}

func CheckUserInfo(user User) error {
	realUser, ok := localStorage[user.Name]
	if !ok {
		return errNotFound
	}

	// do not use it in a prod !!!!
	if user.Password != realUser.Password {
		return errAuth
	}

	return nil
}

func main() {
	user := User{Name: "t.razumov", Password: "qwe"}
	if err := CheckUserInfo(user); err != nil {
		switch err {
		case errAuth:
			fmt.Println("неправильный логин или пароль", err)
		case errNotFound:
			fmt.Println("пользователя не существует. перейти на окно регистрации?", err)
		default:
			fmt.Println(err)
		}
		return
	}

	fmt.Println("ok job1")

	/*
			users := []User{{Name: "t.e.razumov", Password: "qwe"}, {Name: "k.thompson", Password: "love_go"}}
			if err := CheckUsersInfo(users...); err != nil {
				if err == errAuth { // errors.Is(err, errAuth)
					fmt.Println("неправильный логин или пароль", err)
					return
				}

				// var customErr CustomError      // 2
				if err == errNotFound { // errors.Is(err, errNotFound) errors.As(err, &customErr)
					fmt.Println("пользователя не существует. перейти на окно регистрации?", err)
					return
				}

				fmt.Println(err)
				fmt.Println("unwrapped error:", errors.Unwrap(err))

				return
			}

		    fmt.Println("ok job2")
	*/
}

func CheckUsersInfo(users ...User) error {
	var resErr error
	for _, user := range users {
		if err := CheckUserInfo(user); err != nil {
			// resErr = errors.Join(resErr, fmt.Errorf("%s %w", user.Name, err))
			// resErr = fmt.Errorf("%s %w %v", user.Name, err, resErr)

			// %w != %v !!!!!!!
			// return fmt.Errorf("%s %w", user.Name, err)
			// return WrapError{Context: user.Name, Err: err}
			return err
		}
	}

	return resErr
}

// ------------------------------------------

type WrapError struct {
	Context string
	Err     error
}

func (s WrapError) Error() string {
	if s.Err != nil {
		return s.Context + ": " + s.Err.Error()
	}
	return ""
}

func (s WrapError) Unwrap() error {
	return s.Err
}

// ------------------------------------------

type CustomError2 struct {
	Info string
	Code int
}

func (s CustomError2) Error() string {
	return "info: " + s.Info + ", status code: " + strconv.Itoa(s.Code)
}
