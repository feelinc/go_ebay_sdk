package trading

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// GetUserRequest request
type GetUserRequest struct {
	ItemID    string
	UserID    string
	IsSummary bool
}

// Name return request name
func (r GetUserRequest) Name() string {
	return "GetUser"
}

// BodyXML return XML body request
func (r GetUserRequest) BodyXML() string {
	xml := ""

	xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemID)
	xml = xml + fmt.Sprintf("<UserID>%s</UserID>", r.UserID)

	if r.IsSummary {
		xml = xml + "<DetailLevel>ReturnSummary</DetailLevel>"
	} else {
		xml = xml + "<DetailLevel>ReturnAll</DetailLevel>"
	}

	return xml
}

// ParseResponse return parsed response
func (r GetUserRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp GetUserResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// GetUserResponse response
type GetUserResponse struct {
	ebaysdk.EbayResponse
	User User
}

// ResponseErrors return response errors
func (r GetUserResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Error.Items
}

// NewGetUser return new GetUser request
func NewGetUser(itemID string, userID string, isSummary bool) ebaysdk.Request {
	return &GetUserRequest{
		ItemID:    itemID,
		UserID:    userID,
		IsSummary: isSummary,
	}
}
