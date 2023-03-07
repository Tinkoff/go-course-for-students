package main

import (
	"fmt"
	"os"
	"regexp"
)

type Person struct {
	Name    string
	Surname string
}

func Run() (err error) {
	/*
		defer func() {
			fmt.Println("defer 3")
			//if r := recover(); r != nil {
			//	err = fmt.Errorf("panic %v", r)
			//}
		}()
	*/

	/*
		v := 5
		defer func(i int) {
			fmt.Println("defer 2")
			fmt.Println("passed via arg", i)
			//fmt.Println("passed via capture", v)
		}(v)
		v++
	*/

	// defer fmt.Println("defer 1")

	file, err := os.Open("errors/3_panic_and_defer/kek.txt")
	if err != nil {
		return err
	}

	// defer file.Close()

	var p *Person
	p.Surname = "kek"

	var i any = p
	_ = i.(int)

	vec := []int{1, 2, 3}
	vec[5]++

	i, _ = regexp.Compile("(?P<Year>\\d{4}") // MustCompile

	panic("test")

	file.Close()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
		return
	}

	Run2()

	fmt.Println("ok")
}

func Run2() {
	idx := 0
	var idxs []int
	for {
		idx++
		idxs = append(idxs, idx)
		defer func(idx int) {
			if idx%100000 == 0 {
				idxs = nil
				fmt.Println("defer", idx)
			}
		}(idx)

		if idx%100000 == 0 {
			fmt.Println("iter", idx)
		}
	}
}
