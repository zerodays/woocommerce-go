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

type CartAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
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

type CartItemVariation struct {
	Attribute string `json:"attribute"`
	Value     string `json:"value"`
}

type CartItem struct {
	Key              string              `json:"key"`
	ID               int                 `json:"id"`
	Quantity         int                 `json:"quantity"`
	Name             string              `json:"name"`
	Summary          string              `json:"summary"`
	ShortDescription string              `json:"short_description"`
	Description      string              `json:"description"`
	SKU              string              `json:"sku"`
	Images           []CartImage         `json:"images"`
	Totals           CartItemTotals      `json:"totals"`
	Variations       []CartItemVariation `json:"variation"`
}

type CartShippingRateInner struct {
	RateID       string `json:"rate_id"`
	Price        Int    `json:"price"`
	MethodID     string `json:"method_id"`
	Selected     bool   `json:"selected"`
	CurrencyCode string `json:"currency_code"`
}

type CartShippingRate struct {
	PackageID     int                     `json:"package_id"`
	Name          string                  `json:"name"`
	ShippingRates []CartShippingRateInner `json:"shipping_rates"`
}

// Cart holds the cart data.
type Cart struct {
	// CartToken is the cart token returned by woocommerce on cart request.
	CartToken string `json:"cart_token"`

	Coupons         []Coupon           `json:"coupons"`
	Items           []CartItem         `json:"items"`
	ShippingAddress CartAddress        `json:"shipping_address"`
	BillingAddress  CartAddress        `json:"billing_address"`
	Totals          CartTotals         `json:"totals"`
	ShippingRates   []CartShippingRate `json:"shipping_rates"`
}
