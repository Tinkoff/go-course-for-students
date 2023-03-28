package main

import (
	"context"
	"fmt"
	"log"
)

// before make the command go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

//go:generate oapi-codegen -o ./client.go --package=main ./openapi.json
func main() {
	c, err := NewClientWithResponses("https://petstore3.swagger.io/api/v3")
	if err != nil {
		log.Fatal(err)
	}

	status := FindPetsByStatusParamsStatusAvailable
	r, err := c.FindPetsByStatusWithResponse(context.Background(), &FindPetsByStatusParams{Status: &status})
	if err != nil {
		log.Fatal(err)
	}

	for _, pet := range *r.JSON200 {
		fmt.Println(pet.Name)
	}
}
