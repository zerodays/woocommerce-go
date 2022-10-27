package woocommerce

// Tax represents a tax object.
type Tax struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	Rate    Float  `json:"rate"`
}
