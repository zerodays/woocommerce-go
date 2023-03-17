package woocommerce

type Customer struct {
	ID               int        `json:"id,omitempty"`
	DateCreated      string     `json:"date_created,omitempty"`
	DateModified     string     `json:"date_modified,omitempty"`
	Email            string     `json:"email,omitempty"`
	FirstName        string     `json:"first_name,omitempty"`
	LastName         string     `json:"last_name,omitempty"`
	Role             string     `json:"role,omitempty"`
	Username         string     `json:"username,omitempty"`
	Billing          Address    `json:"billing,omitempty"`
	Shipping         Address    `json:"shipping,omitempty"`
	IsPayingCustomer bool       `json:"is_paying_customer,omitempty"`
	AvatarUrl        string     `json:"avatar_url,omitempty"`
	MetaData         []MetaData `json:"meta_data,omitempty"`
}
