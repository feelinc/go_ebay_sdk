package ebaysdk

import (
	"fmt"
)

// Trading object
type Trading struct {
	Config  *Config
	Request Request
}

// Execute the request
func (c *Trading) Execute(req Request) (Response, error) {
	c.Request = req
	return BuildRequest(c.Config, c.BuildRequestHeader(), c.BuildRequestData(), req)
}

// BuildRequestHeader build the request header
func (c *Trading) BuildRequestHeader() map[string]string {
	headers := map[string]string{
		"X-EBAY-API-COMPATIBILITY-LEVEL": fmt.Sprintf("%d", c.Config.Version),
		"X-EBAY-API-DEV-NAME":            c.Config.DevID,
		"X-EBAY-API-APP-NAME":            c.Config.AppID,
		"X-EBAY-API-CERT-NAME":           c.Config.CertID,
		"X-EBAY-API-SITEID":              fmt.Sprintf("%d", c.Config.SiteID),
		"X-EBAY-API-CALL-NAME":           c.Request.Name(),
		"Content-Type":                   "text/xml",
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
func (c *Trading) BuildRequestData() string {
	xml := "<?xml version='1.0' encoding='utf-8'?>"
	xml = xml + fmt.Sprintf("<%sRequest xmlns=\"urn:ebay:apis:eBLBaseComponents\">", c.Request.Name())
	xml = xml + "<RequesterCredentials>"
	xml = xml + fmt.Sprintf("<eBayAuthToken>%s</eBayAuthToken>", c.Config.Token)
	xml = xml + "</RequesterCredentials>"
	xml = xml + c.Request.BodyXML()
	xml = xml + fmt.Sprintf("</%sRequest>", c.Request.Name())

	return xml
}
