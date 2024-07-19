package product

import (
	"testing"
	"web/internal/domain"
	"web/internal/errors"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (prm *ProductRepositoryMock) GetAllProducts() ([]domain.Product, error) {
	args := prm.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (prm *ProductRepositoryMock) GetProductByID(id int) (domain.Product, error) {
	args := prm.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (prm *ProductRepositoryMock) SearchProduct(priceGt float64) ([]domain.Product, error) {
	args := prm.Called(priceGt)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (prm *ProductRepositoryMock) CreateProduct(product domain.Product) error {
	args := prm.Called(product)
	return args.Error(0)
}

func (prm *ProductRepositoryMock) UpdateProduct(id int, product domain.Product) error {
	args := prm.Called(id, product)
	return args.Error(0)
}

func (prm *ProductRepositoryMock) PatchProduct(id int, product domain.ProductRequestDto) error {
	args := prm.Called(id, product)
	return args.Error(0)
}

func (prm *ProductRepositoryMock) DeleteProduct(id int) error {
	args := prm.Called(id)
	return args.Error(0)
}

func (prm *ProductRepositoryMock) GetGroupOfProductsByIds(ids []int) ([]domain.Product, error) {
	args := prm.Called(ids)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func TestProductRepository_GetAllProducts(t *testing.T) {

	t.Run("GetAllProducts", func(t *testing.T) {
		prm := ProductRepositoryMock{}
		prm.On("GetAllProducts").Return([]domain.Product{
			{ID: 1, Name: "Product 1", Price: 100.0},
			{ID: 2, Name: "Product 2", Price: 200.0},
		}, nil)

		products, err := prm.GetAllProducts()
		if err != nil {
			t.Fatalf("error getting all products: %v", err)
		}

		if len(products) != 2 {
			t.Fatalf("expected 2 products, got %d", len(products))
		}

		if products[0].ID != 1 {
			t.Fatalf("expected product ID 1, got %d", products[0].ID)
		}

		if products[1].ID != 2 {
			t.Fatalf("expected product ID 2, got %d", products[1].ID)
		}
	})

	t.Run("GetAllProductsError", func(t *testing.T) {
		prm := ProductRepositoryMock{}
		prm.On("GetAllProducts").Return([]domain.Product{}, errors.ErrProductNotFound)

		_, err := prm.GetAllProducts()
		if err == nil {
			t.Fatalf("error expected but not found")
		}
	})

}
