package main

import (
	"fmt"
	"reflect"
)

type CustomInt int

var myCustomInt CustomInt
var i int

func main() {

	i = 5

	typeI := reflect.TypeOf(i)
	valueI := reflect.ValueOf(i)

	fmt.Printf("type: %v, value: %v\n", typeI, valueI)

	typeX := reflect.TypeOf(myCustomInt)
	valueX := reflect.ValueOf(myCustomInt)
	originalTypeX := valueX.Kind() // может различать только basic types
	fmt.Printf("type: %v, value: %v, original type: %v\n", typeX, valueX, originalTypeX)

	val := valueX.Interface() // можно получить значение
	fmt.Printf("CustomInt value is: %v\n", val)
}
