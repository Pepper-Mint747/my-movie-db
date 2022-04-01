package data

import (
	"fmt"
	"strconv"
)

// Runtime declared as custom type
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format
	jsonValue := fmt.Sprintf("%d mins", r)

	//Use strconv.Quote() function on the string to wrap it in doubles.
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}