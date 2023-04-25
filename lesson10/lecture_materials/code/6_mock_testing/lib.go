package lib

import (
	"errors"
	"sync"
)

type Product struct {
	ID    string
	Name  string
	Price int
}

type ProductRepository struct {
	data  map[string]Product
	mutex sync.Mutex
}

func NewProductRepository() ProductRepository {
	return ProductRepository{
		data:  make(map[string]Product),
		mutex: sync.Mutex{},
	}
}

type ProductRepositoryInterface interface {
	Add(product Product) error
}

func (r *ProductRepository) Add(product Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.data[product.ID] = product
	return nil
}

type ProductService struct {
	repo ProductRepositoryInterface
}

func NewProductService(repo ProductRepositoryInterface) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (s ProductService) Insert(productID string, product Product) error {
	if len(productID) == 0 {
		return errors.New("productID can not be null")
	}

	err := s.repo.Add(Product{
		ID:    productID,
		Name:  product.Name,
		Price: product.Price,
	})
	if err != nil {
		return err
	}

	return nil
}
