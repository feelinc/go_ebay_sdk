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
		errs = append(errs, fmt.Sprintf("%#v", e.Error()))
	}

	return strings.Join(errs, ", ")
}

// HTTPError data
type HTTPError struct {
	statusCode int
	body       []byte
}

func (err HTTPError) Error() string {
	return fmt.Sprintf("%d - %s", err.statusCode, err.body)
}
