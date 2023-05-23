package woocommerce

type ErrorDetails struct {
	Code             string         `json:"code"`
	Message          string         `json:"message"`
	AdditionalErrors []ErrorDetails `json:"additional_errors"`
}

type Error struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Data       struct {
		Status  int                     `json:"status"`
		Params  map[string]string       `json:"params"`
		Details map[string]ErrorDetails `json:"details"`
	}
}

func (e *Error) Error() string {
	return e.Message
}
