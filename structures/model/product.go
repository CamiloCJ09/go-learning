package model

var products []Product

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

func (p Product) Save() {
	products = append(products, p)
}

func (p Product) GetAll() []Product {
	return products
}

func GetById(id int) *Product {
	for _, product := range products {
		if product.Id == id {
			return &product
		}
	}
	return nil
}
