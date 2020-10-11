package request

type PlaceOrder struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	OrderId   string `json:"order_id"`
}
