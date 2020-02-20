package trading

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// AddMemberMessageRTQRequest request
type AddMemberMessageRTQRequest struct {
	RecipientID     string
	Body            string
	EmailCopy       string
	ParentMessageID string
	ItemID          string
}

// Name return request name
func (r AddMemberMessageRTQRequest) Name() string {
	return "AddMemberMessageRTQ"
}

// BodyXML return XML body request
func (r AddMemberMessageRTQRequest) BodyXML() string {
	xml := ""

	if r.ItemID != "" {
		xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemID)
	}

	xml = xml + "<MemberMessage>"
	xml = xml + fmt.Sprintf("<RecipientID>%s</RecipientID>", r.RecipientID)
	xml = xml + fmt.Sprintf("<Body>%s</Body>", r.Body)
	xml = xml + fmt.Sprintf("<EmailCopyToSender>%s</EmailCopyToSender>", r.EmailCopy)

	if r.ParentMessageID != "" {
		xml = xml + fmt.Sprintf("<ParentMessageID>%s</ParentMessageID>", r.ParentMessageID)
	}
	xml = xml + "</MemberMessage>"
	xml = xml + "<WarningLevel>High</WarningLevel>"

	return xml
}

// ParseResponse return parsed response
func (r AddMemberMessageRTQRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp AddMemberMessageRTQResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// AddMemberMessageRTQResponse response
type AddMemberMessageRTQResponse struct {
	ebaysdk.EbayResponse
}

// ResponseErrors return response errors
func (r AddMemberMessageRTQResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Errors
}

// NewAddMemberMessageRTQ return new AddMemberMessageAAQToPartner request
func NewAddMemberMessageRTQ(recipientID string, body string, emailCopy bool,
	parentMessageID *string, itemID *string) ebaysdk.Request {
	req := &AddMemberMessageRTQRequest{
		RecipientID: recipientID,
		Body:        body,
	}

	if emailCopy {
		req.EmailCopy = "true"
	} else {
		req.EmailCopy = "false"
	}

	if parentMessageID != nil {
		req.ParentMessageID = *parentMessageID
	}

	if itemID != nil {
		req.ItemID = *itemID
	}

	return req
}
