package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func RunSomeJob() (bool, error) {
	str := "kek1" // "1"
	var err error

	i, err := strconv.Atoi(str)
	if err != nil {
		return false, err
	}

	path := "kek.txt" // "errors/0_basic/kek.txt"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return false, err
	}

	if _, err = f.Write([]byte(time.Now().String() + "\n")); err != nil {
		return false, err
	}

	if err = f.Close(); err != nil {
		return false, err
	}

	return i == 1, nil
}

func main() {
	ok, err := RunSomeJob()
	if err != nil {
		fmt.Println(err)
		return
	}

	if ok {
		fmt.Println("все ок")
	} else {
		fmt.Println("что-то не ок")
		os.Exit(1)
	}
}
