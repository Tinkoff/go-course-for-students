package main

import (
	"fmt"
	"reflect"
)

type CustomInt int

type SomeData struct {
	Foo string `yaml:"foo,omitempty" json:"foo" custom:"bar"`
	Bar int
	Baz CustomInt
}

func main() {

	d := SomeData{
		Foo: "foo",
		Bar: 17,
		Baz: 42,
	}

	dt := reflect.TypeOf(d)
	fmt.Printf("SomeData number of fields: %d\n", dt.NumField())
	fmt.Printf("SomeData field 0 type: %v\n", dt.Field(0).Type)
	fmt.Printf("SomeData field 0 name: %v\n", dt.Field(0).Name)
	fmt.Printf("SomeData field 0 tag: %v\n", dt.Field(0).Tag)
	fmt.Printf("SomeData field 0 tag yaml value: %v\n", dt.Field(0).Tag.Get("yaml"))
	fmt.Printf("SomeData field 0 tag custom value: %v\n", dt.Field(0).Tag.Get("custom"))

}
