package woocommerce

type Customer struct {
	ID               int        `json:"id"`
	DateCreated      string     `json:"date_created"`
	DateModified     string     `json:"date_modified"`
	Email            string     `json:"email"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	Role             string     `json:"role"`
	Username         string     `json:"username"`
	Billing          Address    `json:"billing"`
	Shipping         Address    `json:"shipping"`
	IsPayingCustomer bool       `json:"is_paying_customer"`
	AvatarUrl        string     `json:"avatar_url"`
	MetaData         []MetaData `json:"meta_data"`
}
