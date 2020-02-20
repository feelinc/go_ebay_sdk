package ebaysdk

import (
	"fmt"
	"strings"
)

// Errors responses
type Errors []ResponseError

func (err Errors) Error() string {
	var errs []string

	for _, e := range err {
		errs = append(errs, fmt.Sprintf("%#v", e))
	}

	return strings.Join(errs, ", ")
}

// HttpError data
type HttpError struct {
	statusCode int
	body       []byte
}

func (err HttpError) Error() string {
	return fmt.Sprintf("%d - %s", err.statusCode, err.body)
}
