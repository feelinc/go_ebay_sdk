package ebaysdk

import (
	"fmt"
	"strings"
	"time"
)

// Response interface
type Response interface {
	Success() bool
	Failure() bool
	Acknowledge() string
	Warning() bool
	ResponseErrors() Errors
}

// EbayResponse data
type EbayResponse struct {
	EBayTime  string         `xml:"EBayTime"`
	Timestamp time.Time      `xml:"Timestamp"`
	Ack       string         `xml:"Ack"`
	Build     string         `xml:"Build"`
	Version   string         `xml:"Version"`
	Error     ResponseErrors `xml:"Errors"`
}

// Success return true if success, otherwise false
func (r EbayResponse) Success() bool {
	return r.Ack == "Success"
}

// Failure return true if failure, otherwise false
func (r EbayResponse) Failure() bool {
	return r.Ack == "Failure"
}

// Acknowledge return acknowledge info
func (r EbayResponse) Acknowledge() string {
	return r.Ack
}

// Warning return true if warning, otherwise false
func (r EbayResponse) Warning() bool {
	return r.Ack == "Warning"
}

// ResponseErrors return response errors
func (r EbayResponse) ResponseErrors() Errors {
	return r.Error.Items
}

// ResponseErrors data
type ResponseErrors struct {
	Items []ResponseError `xml:"Error"`
}

func (err ResponseErrors) Error() string {
	var errs []string

	for _, e := range err.Items {
		errs = append(errs, fmt.Sprintf("%#v", e.Error()))
	}

	return strings.Join(errs, ", ")
}

// ResponseError data
type ResponseError struct {
	Code                string `xml:"Code"`
	ErrorClass          string `xml:"ErrorClass"`
	ShortMessage        string `xml:"ShortMessage"`
	LongMessage         string `xml:"LongMessage"`
	ErrorCode           string `xml:"ErrorCode"`
	SeverityCode        string `xml:"SeverityCode"`
	Severity            string `xml:"Severity"`
	Line                string `xml:"Line"`
	Column              string `xml:"Column"`
	ErrorClassification string `xml:"ErrorClassification"`
}

func (err ResponseError) Error() string {
	if err.Code != "" {
		return fmt.Sprintf("[%s] %s", err.Code, err.ShortMessage)
	}
	return fmt.Sprintf("%s", err.ShortMessage)
}
