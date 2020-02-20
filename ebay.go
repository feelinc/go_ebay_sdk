package ebaysdk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/google/uuid"
)

var version = "1.0.0"
var userAgent = fmt.Sprintf("eBaySDK/%s Go/%s", version, runtime.Version())

var httpSsl = map[bool]string{
	false: "http",
	true:  "https",
}

// SiteIds site IDs
var SiteIds = map[string]int{
	"US":    0,
	"GB":    3,
	"SG":    216,
	"FR":    71,
	"FRBE":  0,
	"FRCA":  0,
	"MY":    207,
	"NL":    146,
	"AT":    16,
	"AU":    15,
	"CH":    193,
	"DE":    77,
	"ENCA":  0,
	"ES":    186,
	"HK":    201,
	"IE":    205,
	"IN":    203,
	"IT":    101,
	"MOTOR": 0,
	"NLBE":  0,
	"PH":    211,
	"PL":    212,
}

// Config data
type Config struct {
	Domain              string
	Debug               bool
	Warnings            bool
	Errors              bool
	URI                 string
	DevID               string
	AppID               string
	CertID              string
	Token               string
	SiteID              int
	Version             int
	HTTPS               bool
	Timeout             int
	TrackingID          string
	TrackingPartnerCode string
	RequestEncoding     string
	ResponseEncoding    string
}

// Request interface
type Request interface {
	Name() string
	BodyXML() string
	ParseResponse(content []byte) (Response, error)
}

// Connection interface
type Connection interface {
	Execute(req Request) (Response, error)
	BuildRequestHeader() map[string]string
	BuildRequestData() string
}

// BuildRequest build the request
func BuildRequest(config *Config, headers map[string]string, body string,
	request Request) (Response, error) {

	uri := buildRequestURI(config.HTTPS, config.Domain, config.URI)

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return EbayResponse{}, err
	}

	for k := range headers {
		req.Header.Add(k, headers[k])
	}

	reqID, _ := uuid.NewRandom()

	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("X-EBAY-SDK-REQUEST-ID", reqID.String())

	client := &http.Client{
		Timeout: time.Duration(time.Duration(config.Timeout) * time.Second),
	}
	resp, err := client.Do(req)
	if err != nil {
		return EbayResponse{}, err
	}

	if resp.StatusCode != 200 {
		httpErr := HttpError{
			statusCode: resp.StatusCode,
		}
		httpErr.body, _ = ioutil.ReadAll(resp.Body)

		return EbayResponse{}, httpErr
	}

	bodyContent, _ := ioutil.ReadAll(resp.Body)

	if config.Debug {
		log.Println(string(bodyContent[:]))
	}

	response, err := request.ParseResponse(bodyContent)
	if err != nil {
		return EbayResponse{}, err
	}

	if response.Failure() {
		return response, Errors(response.ResponseErrors())
	}

	return response, nil
}

func buildRequestURI(https bool, domain string, uri string) string {
	return fmt.Sprintf("%s://%s%s", httpSsl[https], domain, uri)
}
