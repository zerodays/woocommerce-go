package woocommerce

type CouponTotals struct {
	CurrencyCode              string `json:"currency_code"`
	CurrencySymbol            string `json:"currency_symbol"`
	CurrencyMinorUnit         int    `json:"currency_minor_unit"`
	CurrencyDecimalSeparator  string `json:"currency_decimal_separator"`
	CurrencyThousandSeparator string `json:"currency_thousand_separator"`
	CurrencyPrefix            string `json:"currency_prefix"`
	CurrencySuffix            string `json:"currency_suffix"`
	TotalDiscount             Int    `json:"total_discount"`
	TotalDiscountTax          Int    `json:"total_discount_tax"`
}

type Coupon struct {
	Code   string       `json:"code"`
	Totals CouponTotals `json:"totals"`
}

type CartImage struct {
	ID        int    `json:"id"`
	SRC       string `json:"src"`
	Thumbnail string `json:"thumbnail"`
}

type CartItemTotals struct {
	CurrencyCode      string `json:"currency_code"`
	CurrencyMinorUnit int    `json:"currency_minor_unit"`
	LineTotal         Int    `json:"line_total"`
	LineTotalTax      Int    `json:"line_total_tax"`
}

type CartTotals struct {
	CurrencyCode      string `json:"currency_code"`
	CurrencyMinorUnit int    `json:"currency_minor_unit"`
	TotalItems        Int    `json:"total_items"`
	TotalItemsTax     Int    `json:"total_items_tax"`
	TotalFees         Int    `json:"total_fees"`
	TotalFeesTax      Int    `json:"total_fees_tax"`
	TotalDiscount     Int    `json:"total_discount"`
	TotalDiscountTax  Int    `json:"total_discount_tax"`
	TotalShipping     Int    `json:"total_shipping"`
	TotalShippingTax  Int    `json:"total_shipping_tax"`
	TotalPrice        Int    `json:"total_price"`
	TotalTax          Int    `json:"total_tax"`
}

type CartItem struct {
	Key              string         `json:"key"`
	ID               int            `json:"id"`
	Quantity         int            `json:"quantity"`
	Name             string         `json:"name"`
	Summary          string         `json:"summary"`
	ShortDescription string         `json:"short_description"`
	Description      string         `json:"description"`
	SKU              string         `json:"sku"`
	Images           []CartImage    `json:"images"`
	Totals           CartItemTotals `json:"totals"`
}

// Cart holds the cart data.
type Cart struct {
	Coupons []Coupon   `json:"coupons"`
	Items   []CartItem `json:"items"`
	Totals  CartTotals `json:"totals"`
}
