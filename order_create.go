package woocommerce

type OrderCreateBilling struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderCreateShipping struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}

type OrderCreateItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderCreateMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OrderShippingLine struct {
	MethodID string `json:"method_id"`
	// Total is the total price of the shipping line.
	// It is formatted on two decimal with '.' as decimal separator.
	Total string `json:"total"`
}

type OrderCreate struct {
	PaymentMethod      string                `json:"payment_method"`
	PaymentMethodTitle string                `json:"payment_method_title"`
	Currency           string                `json:"currency"`
	SetPaid            bool                  `json:"set_paid"`
	Billing            OrderCreateBilling    `json:"billing"`
	Shipping           OrderCreateShipping   `json:"shipping"`
	Items              []OrderCreateItem     `json:"line_items"`
	MetaData           []OrderCreateMetadata `json:"meta_data"`
	ShippingLines      []OrderShippingLine   `json:"shipping_lines"`
}
