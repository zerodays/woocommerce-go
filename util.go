package woocommerce

import (
	"encoding/json"
	"math"
	"strconv"
	"time"
)

const timeFormat = "2006-01-02T15:04:05"

// Float is a support type that marshals itself to string in JSON.
type Float float64

func (f Float) MarshalJSON() ([]byte, error) {
	val := math.Round(float64(f)*100) / 100 // Round to 2 decimal places
	valStr := strconv.FormatFloat(val, 'f', 2, 64)
	return json.Marshal(valStr)
}

func (f *Float) UnmarshalJSON(bytes []byte) error {
	// Try to unmarshal into float.
	var val float64
	err := json.Unmarshal(bytes, &val)
	if err == nil {
		*f = Float(val)
		return nil
	}

	// Unmarshal to float failed. Try to unmarshal to string and
	// convert the string to float.
	var valStr string
	err = json.Unmarshal(bytes, &valStr)
	if err != nil {
		return err
	}

	// Convert to float
	val, err = strconv.ParseFloat(valStr, 64)
	if err != nil {
		return err
	}

	*f = Float(val)
	return nil
}

// Time is a support type that marshals itself into JSON
// without timezone.
type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	val := t.Format(timeFormat)
	return json.Marshal(val)
}

func (t *Time) UnmarshalJSON(bytes []byte) error {
	var valStr string
	err := json.Unmarshal(bytes, &valStr)
	if err != nil {
		return err
	}

	t.Time, err = time.Parse(timeFormat, valStr)
	return err
}
