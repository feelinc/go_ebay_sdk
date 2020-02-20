package trading

import (
	"encoding/xml"
	"fmt"

	"github.com/feelinc/ebaysdk"
)

// GetFeedbackRequest request
type GetFeedbackRequest struct {
	Type   string
	UserID string
}

// Name return request name
func (r GetFeedbackRequest) Name() string {
	return "GetFeedback"
}

// BodyXML return XML body request
func (r GetFeedbackRequest) BodyXML() string {
	xml := ""

	xml = xml + fmt.Sprintf("<FeedbackType>%s</FeedbackType>", r.Type)
	xml = xml + fmt.Sprintf("<UserID>%s</UserID>", r.UserID)
	xml = xml + "<WarningLevel>High</WarningLevel>"

	return xml
}

// ParseResponse return parsed response
func (r GetFeedbackRequest) ParseResponse(content []byte) (ebaysdk.Response, error) {
	var resp GetFeedbackResponse
	err := xml.Unmarshal(content, &resp)
	return resp, err
}

// GetFeedbackResponse response
type GetFeedbackResponse struct {
	ebaysdk.EbayResponse
	FeedbackSummary FeedbackSummary
	FeedbackScore   int
}

// ResponseErrors return response errors
func (r GetFeedbackResponse) ResponseErrors() ebaysdk.Errors {
	return r.EbayResponse.Errors
}

// NewGetFeedback return new GetFeedback request
func NewGetFeedback(theType string, userID string) ebaysdk.Request {
	return &GetFeedbackRequest{
		Type:   theType,
		UserID: userID,
	}
}
