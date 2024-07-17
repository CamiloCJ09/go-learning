package storage

import (
	"encoding/json"
	"errors"
	"os"
	"web/internal/domain"
	internalErrors "web/internal/errors"
)

//Methods to modify the json file with the products
// file to modify products.json file, to add, delete, update products

type productStorage struct {
	fileName string
}

type ProductStorage interface {
	GetAllProducts() ([]domain.Product, error)
	GetProductByID(id int) (domain.Product, error)
	SearchProduct(priceGt float64) ([]domain.Product, error)
	CreateProduct(product domain.Product) error
	UpdateProduct(id int, product domain.Product) error
	PatchProduct(id int, product domain.ProductRequestDto) error
	DeleteProduct(id int) error
}

func NewProductStorage(fileName string) ProductStorage {
	return &productStorage{
		fileName: fileName,
	}
}

func (ps *productStorage) GetAllProducts() ([]domain.Product, error) {
	products, err := ps.readProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *productStorage) GetProductByID(id int) (domain.Product, error) {

	products, err := ps.readProducts()
	if err != nil {
		return domain.Product{}, err
	}

	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}

	return domain.Product{}, internalErrors.ErrProductNotFound
}

func (ps *productStorage) SearchProduct(priceGt float64) ([]domain.Product, error) {
	var filteredProducts []domain.Product
	products, err := ps.readProducts()
	if err != nil {
		return nil, err
	}
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts, nil
}

func (ps *productStorage) CreateProduct(product domain.Product) error {
	products, err := ps.readProducts()
	if err != nil {
		return err
	}

	product.ID = ps.getNextID()
	products = append(products, product)

	err = ps.writeProducts(products)
	if err != nil {
		return err
	}

	return nil
}

func (pr *productStorage) getNextID() int {
	maxID := 0
	products, err := pr.readProducts()
	if err != nil {
		return 0
	}
	for _, product := range products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}

func (ps *productStorage) UpdateProduct(id int, product domain.Product) error {

	products, err := ps.readProducts()
	if err != nil {
		return err
	}

	for i, p := range products {
		if p.ID == id {
			products[i] = product
			err = ps.writeProducts(products)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.Join(internalErrors.ErrProductNotFound, internalErrors.ErrProductNotUpdated)
}

func (ps *productStorage) PatchProduct(id int, product domain.ProductRequestDto) error {

	products, err := ps.readProducts()
	if err != nil {
		return err
	}

	for i, p := range products {
		if p.ID == id {
			if product.Name != nil {
				products[i].Name = *product.Name
			}
			if product.Quantity != nil {
				products[i].Quantity = *product.Quantity
			}
			if product.CodeValue != nil {
				products[i].CodeValue = *product.CodeValue
			}
			if product.IsPublished != nil {
				products[i].IsPublished = *product.IsPublished
			}
			if product.Expiration != nil {
				products[i].Expiration = *product.Expiration
			}
			if product.Price != nil {
				products[i].Price = *product.Price
			}
			return nil
		}
	}

	return errors.Join(internalErrors.ErrProductNotFound, internalErrors.ErrProductNotPatched)
}

func (ps *productStorage) DeleteProduct(id int) error {

	products, err := ps.readProducts()
	if err != nil {
		return err
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			err = ps.writeProducts(products)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.Join(internalErrors.ErrProductNotFound, internalErrors.ErrProductNotDeleted)
}

func (ps *productStorage) readProducts() ([]domain.Product, error) {

	file, err := os.Open(ps.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []domain.Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *productStorage) writeProducts(products []domain.Product) error {

	file, err := os.Create(ps.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(products)
	if err != nil {
		return err
	}

	return nil
}
