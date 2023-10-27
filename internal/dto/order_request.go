package dto

type Order struct {
	Milk  float64 `json:"milk"`
	Skins int     `json:"skins"`
}

type CustomerOrder struct {
	Customer string `json:"customer"`
	Order    Order  `json:"order"`
}

type OrderResponse struct {
	Milk  float64 `json:"milk,omitempty"`
	Skins int     `json:"skins,omitempty"`
}

type OrderStatus string

const (
	Fulfilled   OrderStatus = "fulfilled"
	Partial     OrderStatus = "partial"
	Unfulfilled OrderStatus = "unfulfilled"
)
