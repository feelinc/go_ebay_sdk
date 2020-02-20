package trading

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// GetItemRequest request
type GetItemRequest struct {
	ItemID      string
	DetailLevel string
}

// Name return request name
func (r GetItemRequest) Name() string {
	return "GetItem"
}

// BodyXML return XML body request
func (r GetItemRequest) BodyXML() string {
	xml := ""

	if r.ItemID != "" {
		xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemID)
	}

	if r.DetailLevel != "" {
		xml = xml + fmt.Sprintf("<DetailLevel>%s</DetailLevel>", r.DetailLevel)
	}

	xml = xml + "<WarningLevel>High</WarningLevel>"

	return xml
}

// ParseResponse return parsed response
func (r GetItemRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp GetItemResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// GetItemResponse response
type GetItemResponse struct {
	ebaysdk.EbayResponse
	Item Item
}

// ResponseErrors return response errors
func (r GetItemResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Errors
}

// NewGetItem return new GetItem request
func NewGetItem(itemID string, detailLevel *string) ebaysdk.Request {
	req := &GetItemRequest{
		ItemID: itemID,
	}

	if detailLevel != nil {
		req.DetailLevel = *detailLevel
	}

	return req
}
