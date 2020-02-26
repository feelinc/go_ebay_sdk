package shopping

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// GetMultipleItemsRequest request
type GetMultipleItemsRequest struct {
	ItemIds         []string
	IncludeSelector string
	MessageID       string
}

// Name return request name
func (r GetMultipleItemsRequest) Name() string {
	return "GetMultipleItems"
}

// BodyXML return XML body request
func (r GetMultipleItemsRequest) BodyXML() string {
	xml := ""

	if len(r.ItemIds) > 0 {
		for k := range r.ItemIds {
			xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemIds[k])
		}
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
func (r GetMultipleItemsRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp GetMultipleItemsResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// GetMultipleItemsResponse response
type GetMultipleItemsResponse struct {
	ebaysdk.EbayResponse
	Item []Item
}

// ResponseErrors return response errors
func (r GetMultipleItemsResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Error.Items
}

// NewGetMultipleItems return new GetMultipleItems request
func NewGetMultipleItems(items []string, includeSelector *string, messageID *string) ebaysdk.Request {
	req := &GetMultipleItemsRequest{
		ItemIds: items,
	}

	if includeSelector != nil {
		req.IncludeSelector = *includeSelector
	}

	if messageID != nil {
		req.MessageID = *messageID
	}

	return req
}
