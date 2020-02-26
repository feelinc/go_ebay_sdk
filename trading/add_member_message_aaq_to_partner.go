package trading

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// AddMemberMessageAAQToPartnerRequest request
type AddMemberMessageAAQToPartnerRequest struct {
	RecipientID     string
	Subject         string
	Body            string
	EmailCopy       string
	ParentMessageID string
	ItemID          string
}

// Name return request name
func (r AddMemberMessageAAQToPartnerRequest) Name() string {
	return "AddMemberMessageAAQToPartner"
}

// BodyXML return XML body request
func (r AddMemberMessageAAQToPartnerRequest) BodyXML() string {
	xml := ""
	xml = xml + fmt.Sprintf("<ItemID>%s</ItemID>", r.ItemID)
	xml = xml + "<MemberMessage>"
	xml = xml + fmt.Sprintf("<RecipientID>%s</RecipientID>", r.RecipientID)
	xml = xml + fmt.Sprintf("<Body>%s</Body>", r.Body)
	xml = xml + fmt.Sprintf("<EmailCopyToSender>%s</EmailCopyToSender>", r.EmailCopy)
	xml = xml + "<QuestionType>CustomizedSubject</QuestionType>"
	xml = xml + fmt.Sprintf("<Subject>%s</Subject>", r.Subject)

	if r.ParentMessageID != "" {
		xml = xml + fmt.Sprintf("<MessageID>%s</MessageID>", r.ParentMessageID)
	}
	xml = xml + "</MemberMessage>"
	xml = xml + "<WarningLevel>High</WarningLevel>"

	return xml
}

// ParseResponse return parsed response
func (r AddMemberMessageAAQToPartnerRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp AddMemberMessageAAQToPartnerResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// AddMemberMessageAAQToPartnerResponse response
type AddMemberMessageAAQToPartnerResponse struct {
	ebaysdk.EbayResponse
}

// ResponseErrors return response errors
func (r AddMemberMessageAAQToPartnerResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Error.Items
}

// NewAddMemberMessageAAQToPartner return new AddMemberMessageAAQToPartner request
func NewAddMemberMessageAAQToPartner(recipientID string, subject string,
	body string, emailCopy bool, itemID string, parentMessageID *string) ebaysdk.Request {
	req := &AddMemberMessageAAQToPartnerRequest{
		RecipientID: recipientID,
		Subject:     subject,
		Body:        body,
		ItemID:      itemID,
	}

	if emailCopy {
		req.EmailCopy = "true"
	} else {
		req.EmailCopy = "false"
	}

	if parentMessageID != nil {
		req.ParentMessageID = *parentMessageID
	}

	return req
}
