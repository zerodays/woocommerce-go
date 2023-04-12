package woocommerce

type Error struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Data       struct {
		Status int               `json:"status"`
		Params map[string]string `json:"params"`
	}
}

func (e *Error) Error() string {
	return e.Message
}
