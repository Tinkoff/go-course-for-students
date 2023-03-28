package main

import "fmt"

type GetRequest struct {
	ID string
}

func (s *GetRequest) SetID(v string) {
	s.ID = v
}

type UpdateRequest struct {
	ID string
}

func (s *UpdateRequest) SetID(v string) {
	s.ID = v
}

func main() {
	getRequest := GetRequest{ID: "id1"}
	fmt.Printf("GetRequest ID: '%s'\n", getRequest.ID)

	updateRequest := UpdateRequest{ID: "id2"}
	fmt.Printf("UpdateRequest ID: '%s'\n", updateRequest.ID)

	//getRequest := Create[GetRequest]("id1")
	//fmt.Printf("GetRequest ID: '%s'\n", getRequest.ID)
	//
	//updateRequest := Create[UpdateRequest]("id2")
	//fmt.Printf("UpdateRequest ID: '%s'\n", updateRequest.ID)
}

//func Create[T Request](elem string) *T {
//	data := T{}
//	data.SetID(elem)
//	return &data
//}

//type Request interface {
//	GetRequest | UpdateRequest
//
//	SetID(string)
//}

//type SetIDPtr[T Request] interface {
//	*T
//
//	SetID(string)
//}

//func Create[T Request, TPtr SetIDPtr[T]](elem string) *T {
//	var data T
//	dataPtr := TPtr(&data)
//	dataPtr.SetID(elem)
//	return &data
//}
