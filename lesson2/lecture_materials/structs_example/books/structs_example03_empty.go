package books

import (
	"fmt"
	"unsafe"
)

type Bookmark struct{}

func DemonstrateEmptyStruct() {
	return
	b1 := Bookmark{}
	fmt.Println("empty struct size:", unsafe.Sizeof(b1))

	array := [100]Bookmark{}
	fmt.Println("empty struct array size:", unsafe.Sizeof(array))

	//b2 := Bookmark{}
	//fmt.Printf("empty struct addresses: %p, %p, %t\n", &b1, &b2, &b1 == &b2)
}
