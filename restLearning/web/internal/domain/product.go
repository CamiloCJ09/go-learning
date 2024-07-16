package domain

type Product struct {
	ID          int     `json:"id" required:"true"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductRequestDto struct {
	Name        *string  `json:"name,omitempty"`
	Quantity    *int     `json:"quantity,omitempty"`
	CodeValue   *string  `json:"code_value,omitempty"`
	IsPublished *bool    `json:"is_published,omitempty"`
	Expiration  *string  `json:"expiration,omitempty"`
	Price       *float64 `json:"price,omitempty"`
}
