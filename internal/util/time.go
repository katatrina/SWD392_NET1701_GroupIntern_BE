package util

import (
	"time"
)

// CustomDate represents a custom date format YYYY-MM-DD
type CustomDate time.Time

// UnmarshalJSON parses a JSON-encoded byte slice into the CustomDate type
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Define a custom format for parsing
	customFormat := "02/01/2006"
	
	// Parse the input data using the custom format
	parsedTime, err := time.Parse(`"`+customFormat+`"`, string(data))
	if err != nil {
		return err
	}
	
	// Assign the parsed time to the CustomDate field
	*cd = CustomDate(parsedTime)
	
	return nil
}

// MarshalJSON converts a CustomDate to a JSON-encoded byte slice
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// Format the time.Time field in the desired format
	formatted := time.Time(cd).Format(`"02/01/2006"`)
	
	// Return the formatted time as a byte slice
	return []byte(formatted), nil
}
