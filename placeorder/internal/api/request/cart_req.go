package request

type Cart struct {
	OrderId  string    `json:"order_id"`
	User     *User     `json:"user"`
	ItemInfo *ItemInfo `json:"item_info"`
}

type User struct {
	UserId  string `json:"user_id"`
	Address string `json:"address"`
}

type ItemInfo struct {
	ItemId    string  `json:"item_id"`
	ItemName  string  `json:"item_name"`
	BasePrice float64 `json:"base_price"`
	Quantity  int     `json:"quantity"`
}