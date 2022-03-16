package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/http"
)

// SignatureHeader is the name of the HTTP header in which the webhook signature
// of the payload is given.
const SignatureHeader = "X-WC-Webhook-Signature"

// CheckPayload checks if the signature of the data
// is correct, using the webhook secret key.
func CheckPayload(data []byte, signature, secret string) (bool, error) {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	hSum := h.Sum(nil)

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("[woocommerce-go]: could not decode signature as base64: %w", err)
	}

	return subtle.ConstantTimeCompare(hSum, signatureBytes) == 1, nil
}

// CheckRequest is a helper function that checks weather the signature
// in the request header for the body data is correct.
func CheckRequest(body []byte, r *http.Request, secret string) (bool, error) {
	signature := r.Header.Get(SignatureHeader)
	return CheckPayload(body, signature, secret)
}
