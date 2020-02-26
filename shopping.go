package ebaysdk

import (
	"fmt"
)

// Shopping object
type Shopping struct {
	Config  *Config
	Request Request
}

// Execute the request
func (c *Shopping) Execute(req Request) (Response, error) {
	c.Request = req
	return BuildRequest(c.Config, c.BuildRequestHeader(), c.BuildRequestData(), req)
}

// BuildRequestHeader build the request header
func (c *Shopping) BuildRequestHeader() map[string]string {
	headers := map[string]string{
		"X-EBAY-API-VERSION":          fmt.Sprintf("%d", c.Config.Version),
		"X-EBAY-API-APP-ID":           c.Config.AppID,
		"X-EBAY-API-SITE-ID":          fmt.Sprintf("%d", c.Config.SiteID),
		"X-EBAY-API-CALL-NAME":        c.Request.Name(),
		"X-EBAY-API-REQUEST-ENCODING": "XML",
		"X-EBAY-API-DETAIL-LEVEL":     "0",
		"Content-Type":                "text/xml",
	}

	if c.Config.TrackingID != "" {
		headers["X-EBAY-API-TRACKING-ID"] = c.Config.TrackingID
	}

	if c.Config.TrackingID != "" {
		headers["X-EBAY-API-TRACKING-PARTNER-CODE"] = c.Config.TrackingPartnerCode
	}

	return headers
}

// BuildRequestData build the request data
func (c *Shopping) BuildRequestData() string {
	xml := "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	xml = xml + fmt.Sprintf("<%sRequest xmlns=\"urn:ebay:apis:eBLBaseComponents\">", c.Request.Name())
	xml = xml + c.Request.BodyXML()
	xml = xml + fmt.Sprintf("</%sRequest>", c.Request.Name())

	return xml
}
