package woocommerce

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusOnHold     OrderStatus = "on-hold"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
	OrderStatusFailed     OrderStatus = "failed"
	OrderStatusTrash      OrderStatus = "trash"
)

// Address is the address used in the order.
// It is used as billing and shipping address.
// Some fields are only present in billing address.
type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`

	// The following fields are only present in the billing address.

	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

// MetaData holds the meta data.
type MetaData struct {
	ID    int         `json:"id"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type OrderTax struct {
	ID       int    `json:"id"`
	RateCode string `json:"rate_code"`
	RateID   String `json:"rate_id"`
	Label    string `json:"label"`
	Compound bool   `json:"compound"`

	// TaxTotal presents tax total not including shipping taxes.
	TaxTotal         Float      `json:"tax_total"`
	ShippingTaxTotal Float      `json:"shipping_tax_total"`
	MetaData         []MetaData `json:"meta_data"`
}

type OrderItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProductID   int    `json:"product_id"`
	VariationID int    `json:"variation_id"`
	Quantity    int    `json:"quantity"`
	TaxClass    string `json:"tax_class"`

	// Subtotal is line subtotal before discounts
	Subtotal Float `json:"subtotal"`
	// SubtotalTax is line subtotal tax before discounts
	SubtotalTax Float `json:"subtotal_tax"`
	// Total is line total after discounts
	Total Float `json:"total"`
	//  TotalTax is line total tax after discounts
	TotalTax Float      `json:"total_tax"`
	Taxes    []OrderTax `json:"taxes"`
	MetaData []MetaData `json:"meta_data"`
	SKU      string     `json:"sku"`
	Price    Float      `json:"price"`
}

type OrderShipping struct {
	ID          int    `json:"id"`
	MethodTitle string `json:"method_title"`
	MethodID    string `json:"method_id"`
	// Total is the line total after discounts.
	Total Float `json:"total"`
	// TotalTax is the line total tax after discounts.
	TotalTax Float      `json:"total_tax"`
	Taxes    []OrderTax `json:"taxes"`
	MetaData []MetaData `json:"meta_data"`
}

type OrderCoupon struct {
	ID          int        `json:"ID"`
	Code        string     `json:"code"`
	Discount    Float      `json:"discount"`
	DiscountTax Float      `json:"discountTax"`
	MetaData    []MetaData `json:"metaData"`
}

type OrderRefund struct {
	ID     int
	Reason string
	// Total is the refund total.
	Total Float
}

// Order is the order object that the API returns.
type Order struct {
	ID         int         `json:"id"`
	ParentID   int         `json:"parent_id"`
	Number     string      `json:"number"`
	OrderKey   string      `json:"order_key"`
	CreatedVia string      `json:"created_via"`
	Version    string      `json:"version"`
	Status     OrderStatus `json:"status"`
	// Currency in 3-letter ISO format.
	Currency     string `json:"currency"`
	DateCreated  Time   `json:"date_created"`
	DateModified Time   `json:"date_modified"`
	// DiscountTotal is the total discount amount for the order
	DiscountTotal string `json:"discount_total"`
	// DiscountTax is the discount tax amount for the order
	DiscountTax string `json:"discount_tax"`
	// ShippingTotal is the shipping amount for the order
	ShippingTotal string `json:"shipping_total"`
	// ShippingTax is the shipping tax amount for the order
	ShippingTax string `json:"shipping_tax"`
	// CartTax is the sum of line item taxes only
	CartTax string `json:"cart_tax"`
	// Total is the grand total
	Total string `json:"total"`
	// TotalTax is the sum of all taxes
	TotalTax string `json:"total_tax"`
	// PricesIncludeTax is true the prices included tax during checkout
	PricesIncludeTax  bool    `json:"prices_include_tax"`
	CustomerID        int     `json:"customer_id"`
	CustomerIPAddress string  `json:"customer_ip_address"`
	CustomerUserAgent string  `json:"customer_user_agent"`
	CustomerNote      string  `json:"customer_note"`
	Billing           Address `json:"billing"`
	Shipping          Address `json:"shipping"`
	// PaymentMethod is the ID of the Payment method
	PaymentMethod      string   `json:"payment_method"`
	PaymentMethodTitle string   `json:"payment_method_title"`
	TransactionID      string   `json:"transaction_id"`
	DatePaid           NullTime `json:"date_paid"`
	DateCompleted      NullTime `json:"date_completed"`
	// CartHash is the MD5 hash of the cart items.
	CartHash      string          `json:"cart_hash"`
	MetaData      []MetaData      `json:"meta_data"`
	LineItems     []OrderItem     `json:"line_items"`
	TaxLines      []OrderTax      `json:"tax_lines"`
	ShippingLines []OrderShipping `json:"shipping_lines"`
	CouponLines   []OrderCoupon   `json:"coupon_lines"`
	Refunds       []OrderRefund   `json:"refunds"`
}
