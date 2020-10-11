package response

type PlaceOrder struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}
