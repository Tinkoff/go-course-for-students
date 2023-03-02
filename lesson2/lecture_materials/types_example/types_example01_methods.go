package main

import "fmt"

type Duration int64

const Millisecond Duration = 1000000

func (d Duration) Milliseconds() int64 { return int64(d) / 1000000 }

func Milliseconds(d Duration) int64 { return int64(d) / 1000000 }

func DemonstrateMethods() {
	return
	ms := Millisecond

	fmt.Println("ms.Milliseconds() = ", ms.Milliseconds())
	fmt.Println("Milliseconds(ms) = ", Milliseconds(ms))
	fmt.Println("ms = ", ms)
}

//
//
//

func (d Duration) Reset() {
	d = 0
}

func DemonstrateMutators() {
	return
	ms := Millisecond
	ms.Reset()
	fmt.Println("ms =", ms)

	//ms2 := &ms
	//ms2.Reset()

	//ms3 := &ms2
	//ms3.Reset()
}
