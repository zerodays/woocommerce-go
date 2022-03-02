package woocommerce

import (
	"net/url"
	"strconv"
)

// Parameters specifies a type that can return GET parameters of the request.
type Parameters interface {
	Values() url.Values
}

// BaseParameters is a wrapper around url.Values type.
type BaseParameters url.Values

func (b BaseParameters) Values() url.Values {
	return url.Values(b)
}

// ParametersWithValue sets the value of the given parameters and returns new parameters.
func ParametersWithValue(p Parameters, key, value string) Parameters {
	v := p.Values()
	v.Set(key, value)
	return BaseParameters(v)
}

// PageParams represents parameters that are used to specify pagination values
// of list queries.
type PageParams struct {
	Page, PerPage int
}

func (p PageParams) Values() url.Values {
	values := url.Values{}
	values.Set("page", strconv.Itoa(p.Page))
	values.Set("per_page", strconv.Itoa(p.PerPage))
	return values
}
