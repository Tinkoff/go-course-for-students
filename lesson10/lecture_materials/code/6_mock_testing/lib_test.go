package lib_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	lib "lib/code/6_mock_testing"
	"lib/code/6_mock_testing/mocks"
	"testing"
)

func TestProductService_Insert(t *testing.T) {
	repo := &mocks.ProductRepositoryInterface{}
	repo.On("Add", mock.AnythingOfType("lib.Product")).
		Return(nil).
		Once()

	service := lib.NewProductService(repo)

	err := service.Insert("2f1afe98-63c4-4f59-bcaf-1df835602bdb", lib.Product{
		Name:  "Macbook",
		Price: 20500,
	})

	assert.Nil(t, err)
}
