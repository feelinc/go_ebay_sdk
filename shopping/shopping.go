package shopping

import (
	"github.com/feelinc/ebaysdk"
)

const (
	domain  = "open.api.ebay.com"
	uri     = "/shopping"
	siteID  = 0
	version = 967
	timeout = 20
)

// OptFunc is a function that configures a shopping.
type OptFunc func(*ebaysdk.Shopping) error

// SetDomain API endpoint (default: open.api.ebay.com).
func SetDomain(domain string) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.Domain = domain
		return nil
	}
}

// SetDebug debugging status (default: false).
func SetDebug(is bool) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.Debug = is
		return nil
	}
}

// SetWarnings warnings status (default: true).
func SetWarnings(is bool) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.Warnings = is
		return nil
	}
}

// SetErrors errors status (default: true).
func SetErrors(is bool) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.Errors = is
		return nil
	}
}

// SetURI API endpoint uri (default: /shopping).
func SetURI(uri string) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.URI = uri
		return nil
	}
}

// SetAppID eBay application id.
func SetAppID(id string) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.AppID = id
		return nil
	}
}

// SetSiteID eBay country site id (default: 0 (US)).
func SetSiteID(id int) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.SiteID = id
		return nil
	}
}

// SetVersion version number (default: 967).
func SetVersion(version int) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.Version = version
		return nil
	}
}

// SetHTTPS execute of https (default: True).
func SetHTTPS(is bool) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.HTTPS = is
		return nil
	}
}

// SetTimeout HTTP request timeout (default: 20).
func SetTimeout(second int) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.Timeout = second
		return nil
	}
}

// SetTrackingID ID to identify you to your tracking partner.
func SetTrackingID(id string) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.TrackingID = id
		return nil
	}
}

// SetTrackingPartnerCode third party who is your tracking partner.
func SetTrackingPartnerCode(code string) OptFunc {
	return func(c *ebaysdk.Shopping) error {
		c.Config.TrackingPartnerCode = code
		return nil
	}
}

// Item response data
type Item struct {
	ItemID        string
	Title         string
	Quantity      int
	QuantitySold  int
	EndTime       string
	ListingStatus string
}

// IsSelingActive return true if active, otherwise false
func (i Item) IsSelingActive() bool {
	return i.ListingStatus == "Active"
}

// NewConnection return new shopping API connection
func NewConnection(options ...OptFunc) ebaysdk.Connection {
	cfg := &ebaysdk.Config{
		Domain:              domain,
		Debug:               false,
		Warnings:            true,
		Errors:              true,
		URI:                 uri,
		SiteID:              siteID,
		Version:             version,
		HTTPS:               true,
		Timeout:             timeout,
		TrackingID:          "",
		TrackingPartnerCode: "",
	}
	c := &ebaysdk.Shopping{
		Config: cfg,
	}

	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil
		}
	}

	return c
}
