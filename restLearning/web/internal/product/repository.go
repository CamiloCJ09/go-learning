package product

import (
	"web/cmd/server/storage"
	"web/internal/domain"
)

type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	CreateProduct(product domain.Product) error
	UpdateProduct(id int, product domain.Product) error
	PatchProduct(id int, product domain.ProductRequestDto) error
	DeleteProduct(id int) error
}

// productRepository is a concrete implementation of ProductRepository
type productRepository struct {
	store storage.ProductStorage
}

// NewProductRepository creates a new ProductRepository with the necessary dependencies
func NewProductRepository(filename string) (ProductRepository, error) {
	repo := &productRepository{}
	repo.store = storage.NewProductStorage(filename)
	return repo, nil
}

func (pr *productRepository) GetAllProducts() ([]domain.Product, error) {

	return pr.store.GetAllProducts()
}

func (pr *productRepository) GetProductByID(id int) (domain.Product, error) {

	return pr.store.GetProductByID(id)
}

func (pr *productRepository) CreateProduct(product domain.Product) error {

	return pr.store.CreateProduct(product)
}

func (pr *productRepository) UpdateProduct(id int, product domain.Product) error {

	return pr.store.UpdateProduct(id, product)
}

func (pr *productRepository) PatchProduct(id int, product domain.ProductRequestDto) error {

	return pr.store.PatchProduct(id, product)
}

func (pr *productRepository) DeleteProduct(id int) error {

	return pr.store.DeleteProduct(id)
}
