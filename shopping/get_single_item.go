package shopping

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// GetSingleItemRequest request
type GetSingleItemRequest struct {
	ItemID          string
	VariationSKU    string
	IncludeSelector string
	MessageID       string
}

// Name return request name
func (r GetSingleItemRequest) Name() string {
	return "GetSingleItem"
}

// BodyXML return XML body request
func (r GetSingleItemRequest) BodyXML() string {
	xml := ""

	if r.ItemID != "" {
		xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemID)
	}

	if r.VariationSKU != "" {
		xml = xml + fmt.Sprintf("<VariationSKU>%s</VariationSKU>", r.VariationSKU)
	}

	if r.IncludeSelector != "" {
		xml = xml + fmt.Sprintf("<IncludeSelector>%s</IncludeSelector>", r.IncludeSelector)
	}

	if r.MessageID != "" {
		xml = xml + fmt.Sprintf("<MessageID>%s</MessageID>", r.MessageID)
	}

	return xml
}

// ParseResponse return parsed response
func (r GetSingleItemRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp GetSingleItemResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// GetSingleItemResponse response
type GetSingleItemResponse struct {
	ebaysdk.EbayResponse
	Item Item
}

// ResponseErrors return response errors
func (r GetSingleItemResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Error.Items
}

// NewGetSingleItem return new GetSingleItem request
func NewGetSingleItem(itemID string, variationSKU *string,
	includeSelector *string, messageID *string) ebaysdk.Request {
	req := &GetSingleItemRequest{
		ItemID: itemID,
	}

	if variationSKU != nil {
		req.VariationSKU = *variationSKU
	}

	if includeSelector != nil {
		req.IncludeSelector = *includeSelector
	}

	if messageID != nil {
		req.MessageID = *messageID
	}

	return req
}
