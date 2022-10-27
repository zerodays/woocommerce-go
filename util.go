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

type NullTime struct {
	Time  Time
	Valid bool
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(t.Time)
}

func (t *NullTime) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		t.Valid = false
		return nil
	}

	if err := json.Unmarshal(bytes, &t.Time); err != nil {
		return err
	}

	t.Valid = true
	return nil
}

type Int int64

func (i Int) MarshalJSON() ([]byte, error) {
	val := strconv.FormatInt(int64(i), 10)
	return json.Marshal(val)
}

func (i *Int) UnmarshalJSON(bytes []byte) error {
	var val int64
	err := json.Unmarshal(bytes, &val)
	if err == nil {
		*i = Int(val)
		return nil
	}

	// Unmarshal to int64 failed. Try to unmarshal to string and
	// convert the string to int64.
	var valStr string
	err = json.Unmarshal(bytes, &valStr)
	if err != nil {
		return err
	}

	// Convert to int64
	val, err = strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return err
	}

	*i = Int(val)
	return nil
}

// String is a string that is sometimes a number (thanks woocommerce).
type String string

func (s *String) UnmarshalJSON(bytes []byte) error {
	var valStr string
	err := json.Unmarshal(bytes, &valStr)
	if err == nil {
		*s = String(valStr)
		return nil
	}

	var valNum int
	err = json.Unmarshal(bytes, &valNum)
	if err == nil {
		*s = String(strconv.Itoa(valNum))
		return nil
	} else {
		return err
	}
}
