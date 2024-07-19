package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"web/internal/domain"

	"github.com/stretchr/testify/mock"
)

type ProductHandlerMock struct {
	mock.Mock
}

var req *http.Request = httptest.NewRequest("GET", "/api/v1/products", nil)
var w *httptest.ResponseRecorder = httptest.NewRecorder()

type ProductServiceMock struct {
	mock.Mock
}

func (prm *ProductServiceMock) GetAllProducts() ([]domain.Product, error) {
	args := prm.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (prm *ProductServiceMock) GetProductByID(id int) (domain.Product, error) {
	args := prm.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (prm *ProductServiceMock) CreateProduct(p domain.Product) error {
	args := prm.Called(p)
	return args.Error(0)
}

func (prm *ProductServiceMock) UpdateProduct(i int, p domain.Product) error {
	args := prm.Called(i, p)
	return args.Error(0)
}

func (prm *ProductServiceMock) DeleteProduct(id int) error {
	args := prm.Called(id)
	return args.Error(0)
}

func (prm *ProductServiceMock) GetGroupOfProductsByIds(ids []int) ([]domain.Product, error) {
	args := prm.Called(ids)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (prm *ProductServiceMock) PatchProduct(id int, p domain.ProductRequestDto) error {
	args := prm.Called(id, p)
	return args.Error(0)
}

func (prm *ProductServiceMock) GetConsumerPrice(ids []int) ([]domain.ConsumerPrice, error) {
	args := prm.Called(ids)
	return args.Get(0).([]domain.ConsumerPrice), args.Error(1)
}

func (prm *ProductServiceMock) SearchProduct(priceGt float64) ([]domain.Product, error) {
	args := prm.Called(priceGt)
	return args.Get(0).([]domain.Product), args.Error(1)
}

func NewProductService() *ProductServiceMock {
	return &ProductServiceMock{}
}

func NewProductHandlerMock() *ProductHandlerMock {
	return &ProductHandlerMock{}
}

func TestGetAllProducts(t *testing.T) {

	service := NewProductService()
	handler := NewProductHandler(service)

	t.Run("Get all products", func(t *testing.T) {
		service.On("GetAllProducts").Return([]domain.Product{}, nil)
		handler.GetAllProducts(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", w.Code)
		}

		service.AssertExpectations(t)
	})

}
