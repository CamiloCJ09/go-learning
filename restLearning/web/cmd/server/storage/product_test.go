package storage

import (
	"testing"
	"web/internal/domain"
	"web/internal/errors"

	"github.com/stretchr/testify/mock"
)

type ProductStorageMock struct {
	mock.Mock
}

func (psm *ProductStorageMock) GetAllProducts() ([]domain.Product, error) {
	args := psm.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (psm *ProductStorageMock) GetProductByID(id int) (domain.Product, error) {
	args := psm.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (psm *ProductStorageMock) SearchProduct(priceGt float64) ([]domain.Product, error) {
	args := psm.Called(priceGt)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (psm *ProductStorageMock) CreateProduct(product domain.Product) error {
	args := psm.Called(product)
	return args.Error(0)
}

func (psm *ProductStorageMock) UpdateProduct(id int, product domain.Product) error {
	args := psm.Called(id, product)
	return args.Error(0)
}

func (psm *ProductStorageMock) PatchProduct(id int, product domain.ProductRequestDto) error {
	args := psm.Called(id, product)
	return args.Error(0)
}

func (psm *ProductStorageMock) DeleteProduct(id int) error {
	args := psm.Called(id)
	return args.Error(0)
}

func (psm *ProductStorageMock) GetGroupOfProductsByIds(ids []int) ([]domain.Product, error) {
	args := psm.Called(ids)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func TestProductStorage_GetAllProducts(t *testing.T) {

	t.Run("GetAllProducts", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("GetAllProducts").Return([]domain.Product{
			{ID: 1, Name: "Product 1", Price: 100.0},
			{ID: 2, Name: "Product 2", Price: 200.0},
			{ID: 3, Name: "Product 3", Price: 300.0},
			{ID: 4, Name: "Product 4", Price: 400.0},
		}, nil)

		_, err := psm.GetAllProducts()
		if err != nil {
			t.Errorf("Error getting all products: %v", err)
		}

		psm.AssertExpectations(t)
	})

	t.Run("GetAllProductsError", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("GetAllProducts").Return([]domain.Product{}, errors.ErrProductNotFound)

		_, err := psm.GetAllProducts()
		if err == nil {
			t.Errorf("Error expected but not found")
		}

		psm.AssertExpectations(t)
	})

}

func TestProductStorage_GetProductByID(t *testing.T) {

	t.Run("GetProductByID", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("GetProductByID", 1).Return(domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		}, nil)

		_, err := psm.GetProductByID(1)
		if err != nil {
			t.Errorf("Error getting product by ID: %v", err)
		}

		psm.AssertExpectations(t)
	})

	t.Run("GetProductByIDError", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("GetProductByID", 1).Return(domain.Product{}, errors.ErrProductNotFound)

		_, err := psm.GetProductByID(1)
		if err == nil {
			t.Errorf("Error expected but not found")
		}
	})
}

func TestProductStorage_SearchProduct(t *testing.T) {

	t.Run("SearchProduct", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("SearchProduct", 100.0).Return([]domain.Product{
			{ID: 1, Name: "Product 1", Price: 100.0},
			{ID: 2, Name: "Product 2", Price: 200.0},
			{ID: 3, Name: "Product 3", Price: 300.0},
			{ID: 4, Name: "Product 4", Price: 400.0},
		}, nil)

		_, err := psm.SearchProduct(100.0)
		if err != nil {
			t.Errorf("Error searching product: %v", err)
		}

		psm.AssertExpectations(t)
	})

	t.Run("SearchProductError", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("SearchProduct", 100.0).Return([]domain.Product{}, errors.ErrProductNotFound)

		_, err := psm.SearchProduct(100.0)
		if err == nil {
			t.Errorf("Error expected but not found")
		}
	})
}

func TestProductStorage_CreateProduct(t *testing.T) {

	t.Run("CreateProduct", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("CreateProduct", domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		}).Return(nil)

		err := psm.CreateProduct(domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		})
		if err != nil {
			t.Errorf("Error creating product: %v", err)
		}

		psm.AssertExpectations(t)
	})

	t.Run("CreateProductError", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("CreateProduct", domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		}).Return(errors.ErrProductNotFound)

		err := psm.CreateProduct(domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		})
		if err == nil {
			t.Errorf("Error expected but not found")
		}
	})
}

func TestProductStorage_UpdateProduct(t *testing.T) {

	t.Run("UpdateProduct", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("UpdateProduct", 1, domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		}).Return(nil)

		err := psm.UpdateProduct(1, domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		})
		if err != nil {
			t.Errorf("Error updating product: %v", err)
		}

		psm.AssertExpectations(t)
	})

	t.Run("UpdateProductError", func(t *testing.T) {
		psm := ProductStorageMock{}
		psm.On("UpdateProduct", 1, domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		}).Return(errors.ErrProductNotFound)

		err := psm.UpdateProduct(1, domain.Product{
			ID:    1,
			Name:  "Product 1",
			Price: 100.0,
		})
		if err == nil {
			t.Errorf("Error expected but not found")
		}
	})
}
