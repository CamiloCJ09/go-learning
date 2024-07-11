package model

type Product interface {
	GetPrice() float64
}

type BaseProduct struct {
	Price float64 `json:"price"`
}

type Small struct {
	Product BaseProduct `json:"product"`
}
type Medium struct {
	Product BaseProduct `json:"product"`
}
type Large struct {
	Product BaseProduct `json:"product"`
}

func CreateProduct(price float64, productType string) Product {
	switch productType {
	case "Small":
		return &Small{Product: BaseProduct{Price: price}}
	case "Medium":
		return &Medium{Product: BaseProduct{Price: price}}
	case "Large":
		return &Large{Product: BaseProduct{Price: price}}
	default:
		return nil
	}
}

func (s Small) GetPrice() float64 {
	return s.Product.Price
}

func (m Medium) GetPrice() float64 {
	return m.Product.Price * 1.03
}

func (l Large) GetPrice() float64 {
	return (l.Product.Price * 1.06) + 2500
}
