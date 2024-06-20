package util

import (
	"time"
)

// CustomDate type that embeds time.Time
type CustomDate struct {
	time.Time
}

// UnmarshalJSON method to parse the custom time format
func (ct *CustomDate) UnmarshalJSON(b []byte) error {
	// Remove the quotes from the JSON string
	str := string(b)
	if str == "null" {
		ct.Time = time.Time{}
		return nil
	}
	str = str[1 : len(str)-1] // Strip quotes
	
	// Define the expected format
	const layout = "2006-01-02"
	t, err := time.Parse(layout, str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}
