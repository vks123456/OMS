package response

type Cart struct {
	Error   string `json:"error"`
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}
