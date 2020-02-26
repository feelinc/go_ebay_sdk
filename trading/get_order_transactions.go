package trading

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// ItemTransactionID struct
type ItemTransactionID struct {
	ItemID        string
	TransactionID string
}

// GetOrderTransactionsRequest request
type GetOrderTransactionsRequest struct {
	ItemIDs     []ItemTransactionID
	DetailLevel string
}

// Name return request name
func (r GetOrderTransactionsRequest) Name() string {
	return "GetOrderTransactions"
}

// BodyXML return XML body request
func (r GetOrderTransactionsRequest) BodyXML() string {
	xml := "<ItemTransactionIDArray>"

	if len(r.ItemIDs) > 0 {
		for _, val := range r.ItemIDs {
			xml = xml + "<ItemTransactionID>"
			xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", val.ItemID)
			xml = xml + fmt.Sprintf("<TransactionID>%s</TransactionID>", val.TransactionID)
			xml = xml + "</ItemTransactionID>"
		}
	}

	xml = xml + "</ItemTransactionIDArray>"

	if r.DetailLevel != "" {
		xml = xml + fmt.Sprintf("<DetailLevel>%s</DetailLevel>", r.DetailLevel)
	}

	xml = xml + "<WarningLevel>High</WarningLevel>"

	return xml
}

// ParseResponse return parsed response
func (r GetOrderTransactionsRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp GetOrderTransactionsResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// GetOrderTransactionsResponse response
type GetOrderTransactionsResponse struct {
	ebaysdk.EbayResponse
	OrderArray OrderArray
}

// ResponseErrors return response errors
func (r GetOrderTransactionsResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Error.Items
}

// NewGetOrderTransactions return new GetOrderTransactions request
func NewGetOrderTransactions(itemIDs []ItemTransactionID, detailLevel *string) ebaysdk.Request {
	req := &GetOrderTransactionsRequest{
		ItemIDs: itemIDs,
	}

	if detailLevel != nil {
		req.DetailLevel = *detailLevel
	}

	return req
}
