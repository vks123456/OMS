package request

type AddQuantity struct {
	ProductId     string `json:"product_id"`
	AddQuantity int   `json:"add_quantity"`
}
