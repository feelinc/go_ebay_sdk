package trading

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/feelinc/ebaysdk"
)

// CompleteSaleRequest request
type CompleteSaleRequest struct {
	ItemIDs       []string
	OrderID       string
	TransactionID string
	ShippedTime   *time.Time
}

// CompleteSaleOption is a function that configures a request
type CompleteSaleOption func(*CompleteSaleRequest) error

// SetItemIDs can be used to specify the item IDs slice
func SetItemIDs(ids []string) CompleteSaleOption {
	return func(r *CompleteSaleRequest) error {
		r.ItemIDs = ids
		return nil
	}
}

// SetOrderID can be used to specify the order ID
func SetOrderID(id string) CompleteSaleOption {
	return func(r *CompleteSaleRequest) error {
		r.OrderID = id
		return nil
	}
}

// SetTransactionID can be used to specify the transaction ID
func SetTransactionID(id string) CompleteSaleOption {
	return func(r *CompleteSaleRequest) error {
		r.TransactionID = id
		return nil
	}
}

// SetShippedTime can be used to specify the shipped time
func SetShippedTime(time *time.Time) CompleteSaleOption {
	return func(r *CompleteSaleRequest) error {
		r.ShippedTime = time
		return nil
	}
}

// Name return request name
func (r CompleteSaleRequest) Name() string {
	return "CompleteSale"
}

// BodyXML return XML body request
func (r CompleteSaleRequest) BodyXML() string {
	xml := ""

	if len(r.ItemIDs) > 0 {
		for k := range r.ItemIDs {
			xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemIDs[k])
		}
	}

	if r.ShippedTime != nil {
		xml = xml + "<Shipment>"
		xml = xml + fmt.Sprintf("<ShippedTime>%s</ShippedTime>", r.ShippedTime.Format(time.RFC3339))
		xml = xml + "</Shipment>"
		xml = xml + "<Shipped>true</Shipped>"
	}

	if r.OrderID != "" {
		xml = xml + fmt.Sprintf("<OrderID>%s</OrderID>", r.OrderID)
	}

	if r.TransactionID != "" {
		xml = xml + fmt.Sprintf("<TransactionID>%s</TransactionID>", r.TransactionID)
	}

	xml = xml + "<WarningLevel>High</WarningLevel>"

	return xml
}

// ParseResponse return parsed response
func (r CompleteSaleRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp CompleteSaleResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// CompleteSaleResponse response
type CompleteSaleResponse struct {
	ebaysdk.EbayResponse
}

// ResponseErrors return response errors
func (r CompleteSaleResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Errors
}

// NewCompleteSale return new AddMemberMessageAAQToPartner request
func NewCompleteSale(options ...CompleteSaleOption) ebaysdk.Request {
	req := &CompleteSaleRequest{}

	// Run the options on it
	for _, option := range options {
		if err := option(req); err != nil {
			return nil
		}
	}

	return req
}
