package model

type (
	RequestCheckout struct {
		Quantity    int    `json:"quantity"`
		SkuProduct  string `json:"sku_product"`
		ProductName string `json:"product_name"`
	}

	ResponsesCheckout struct {
		Product    ResponseCheckout
		TotalPrice float64
	}

	ResponseCheckout struct {
		Quantity    int    `json:"quantity"`
		SkuProduct  string `json:"sku_product"`
		ProductName string `json:"product_name"`
		PromoItem string `json:"promo_item"`
	}
)
