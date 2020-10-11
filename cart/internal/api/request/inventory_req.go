package request

type GetQuantity struct {
	ProductId string `json:"product_id"`
	FromCache bool   `json:"from_cache"`
}
