package ebaysdk

import "time"

// Response interface
type Response interface {
	Failure() bool
	Warning() bool
	ResponseErrors() Errors
}

// EbayResponse data
type EbayResponse struct {
	Timestamp time.Time
	Ack       string
	Build     string
	Version   string
	Errors    []ResponseError
}

// Failure return true if failure, otherwise false
func (r EbayResponse) Failure() bool {
	return r.Ack == "Failure"
}

// Warning return true if warning, otherwise false
func (r EbayResponse) Warning() bool {
	return r.Ack == "Warning"
}

// ResponseErrors return response errors
func (r EbayResponse) ResponseErrors() Errors {
	return r.Errors
}

// ResponseError data
type ResponseError struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           string
	SeverityCode        string
	ErrorClassification string
}
