package trading

import (
	"time"

	"github.com/feelinc/ebaysdk"
)

const (
	domain  = "api.ebay.com"
	uri     = "/ws/api.dll"
	siteID  = 0
	version = 1081
	timeout = 20
)

// OptFunc is a function that configures a trading.
type OptFunc func(*ebaysdk.Trading) error

// SetDomain API endpoint (default: api.ebay.com).
func SetDomain(domain string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Domain = domain
		return nil
	}
}

// SetDebug debugging status (default: false).
func SetDebug(is bool) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Debug = is
		return nil
	}
}

// SetWarnings warnings status (default: true).
func SetWarnings(is bool) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Warnings = is
		return nil
	}
}

// SetErrors errors status (default: true).
func SetErrors(is bool) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Errors = is
		return nil
	}
}

// SetURI API endpoint uri (default: /ws/api.dll).
func SetURI(uri string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.URI = uri
		return nil
	}
}

// SetDevID eBay developer id.
func SetDevID(id string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.DevID = id
		return nil
	}
}

// SetAppID eBay application id.
func SetAppID(id string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.AppID = id
		return nil
	}
}

// SetCertID eBay certificate id.
func SetCertID(id string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.CertID = id
		return nil
	}
}

// SetToken eBay authorization token.
func SetToken(token string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Token = token
		return nil
	}
}

// SetSiteID eBay country site id (default: 0 (US)).
func SetSiteID(id int) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.SiteID = id
		return nil
	}
}

// SetVersion version number (default: 967).
func SetVersion(version int) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Version = version
		return nil
	}
}

// SetHTTPS execute of https (default: True).
func SetHTTPS(is bool) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.HTTPS = is
		return nil
	}
}

// SetTimeout HTTP request timeout (default: 20).
func SetTimeout(second int) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.Timeout = second
		return nil
	}
}

// SetTrackingID ID to identify you to your tracking partner.
func SetTrackingID(id string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.TrackingID = id
		return nil
	}
}

// SetTrackingPartnerCode third party who is your tracking partner.
func SetTrackingPartnerCode(code string) OptFunc {
	return func(c *ebaysdk.Trading) error {
		c.Config.TrackingPartnerCode = code
		return nil
	}
}

// Item response data
type Item struct {
	ItemID         string
	Title          string
	SellingStatus  SellingStatus
	Currency       string
	Quantity       int
	ListingDetails ListingDetails
}

// IsCheckoutEnabled return true if checkout enabled, otherwise false
func (i Item) IsCheckoutEnabled() bool {
	return i.ListingDetails.CheckoutEnabled == "true"
}

// IsSelingActive return true if selling is active, otherwise false
func (i Item) IsSelingActive() bool {
	return i.SellingStatus.ListingStatus == "Active"
}

// SellingStatus response data
type SellingStatus struct {
	CurrentPrice  float32
	QuantitySold  int
	ListingStatus string
}

// ListingDetails response data
type ListingDetails struct {
	CheckoutEnabled string
	EndTime         string
	ListingStatus   string
}

// OrderArray data
type OrderArray struct {
	Order []Order
}

// Order response data
type Order struct {
	OrderID          string
	OrderStatus      string
	AdjustmentAmount float64
	AmountPaid       float64
	AmountSaved      float64
	CheckoutStatus   CheckoutStatus

	CreatedTime    time.Time
	PaymentMethods string
	SellerEmail    string

	ShippingAddress Address

	Subtotal float64
	Total    float64

	BuyerUserID                         string
	PaidTime                            time.Time
	ShippedTime                         time.Time
	IntegratedMerchantCreditCardEnabled bool
	PaymentHoldStatus                   string
	IsMultiLegShipping                  bool
	SellerUserID                        string
	CancelStatus                        string
	ExtendedOrderID                     string
	ContainseBayPlusTransaction         bool
}

// CheckoutStatus data
type CheckoutStatus struct {
	eBayPaymentStatus                   string
	LastModifiedTime                    time.Time
	PaymentMethod                       string
	Status                              string
	IntegratedMerchantCreditCardEnabled bool
	PaymentInstrument                   string
}

// Address data
type Address struct {
	Name              string
	Street1           string
	Street2           string
	CityName          string
	StateOrProvince   string
	Country           string
	CountryName       string
	Phone             string
	PostalCode        string
	AddressID         string
	AddressOwner      string
	ExternalAddressID string
}

// FeedbackSummary data
type FeedbackSummary struct {
	BidRetractionFeedbackPeriodArray      BidRetractionFeedbackPeriodArray
	NegativeFeedbackPeriodArray           NegativeFeedbackPeriodArray
	NeutralFeedbackPeriodArray            NeutralFeedbackPeriodArray
	PositiveFeedbackPeriodArray           PositiveFeedbackPeriodArray
	TotalFeedbackPeriodArray              TotalFeedbackPeriodArray
	NeutralCommentCountFromSuspendedUsers int
	UniqueNegativeFeedbackCount           int
	UniquePositiveFeedbackCount           int
	UniqueNeutralFeedbackCount            int
	BuyerRoleMetrics                      BuyerRoleMetrics
}

// BidRetractionFeedbackPeriodArray data
type BidRetractionFeedbackPeriodArray struct {
	FeedbackPeriod []FeedbackPeriod
}

// NegativeFeedbackPeriodArray data
type NegativeFeedbackPeriodArray struct {
	FeedbackPeriod []FeedbackPeriod
}

// NeutralFeedbackPeriodArray data
type NeutralFeedbackPeriodArray struct {
	FeedbackPeriod []FeedbackPeriod
}

// PositiveFeedbackPeriodArray data
type PositiveFeedbackPeriodArray struct {
	FeedbackPeriod []FeedbackPeriod
}

// TotalFeedbackPeriodArray data
type TotalFeedbackPeriodArray struct {
	FeedbackPeriod []FeedbackPeriod
}

// FeedbackPeriod data
type FeedbackPeriod struct {
	PeriodInDays int
	Count        int
}

// BuyerRoleMetrics data
type BuyerRoleMetrics struct {
	PositiveFeedbackLeftCount int
	NegativeFeedbackLeftCount int
	NeutralFeedbackLeftCount  int
	FeedbackLeftPercent       float32
}

// User data
type User struct {
	UserID                      string
	UserIDChanged               bool
	Status                      string
	Email                       string
	FeedbackScore               int
	UniqueNegativeFeedbackCount int
	UniquePositiveFeedbackCount int
	PositiveFeedbackPercent     float32
	FeedbackPrivate             bool
	FeedbackRatingStar          string
	IDVerified                  bool
	NewUser                     bool
	eBayGoodStanding            bool
	RegistrationDate            time.Time
	RegistrationAddress         Address
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
	c := &ebaysdk.Trading{
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
