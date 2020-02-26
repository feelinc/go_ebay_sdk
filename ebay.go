package ebaysdk

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
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

	if config.Debug {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Println(err)
		}

		log.Println(string(requestDump))
	}

	client := &http.Client{
		Timeout: time.Duration(time.Duration(config.Timeout) * time.Second),
	}
	resp, err := client.Do(req)
	if err != nil {
		return EbayResponse{}, err
	}

	if resp.StatusCode != 200 {
		httpErr := HTTPError{
			statusCode: resp.StatusCode,
		}
		httpErr.body, _ = ioutil.ReadAll(resp.Body)

		return EbayResponse{}, httpErr
	}

	bodyContent, _ := ioutil.ReadAll(resp.Body)

	// sometimes eBay return unexpected response, which is in binary
	// so we need to decode it first to get the real response
	re := regexp.MustCompile(`<ns:binary.*?>(.*)</ns:binary>`)
	rematch := re.FindSubmatch(bodyContent)
	if len(rematch) >= 1 {
		bodyContent, err := base64.StdEncoding.DecodeString(string(rematch[1]))
		if err != nil {
			log.Println("Error decoding binary base64: ", string(rematch[1]))
		}

		if config.Debug {
			log.Println(string(bodyContent))
		}
	} else {
		if config.Debug {
			log.Println(string(bodyContent[:]))
		}
	}

	response, err := request.ParseResponse(bodyContent)
	if err != nil {
		return EbayResponse{}, err
	}

	if response.Failure() || response.Acknowledge() == "" {
		errs := response.ResponseErrors()
		if len(errs) > 0 {
			return response, errs[0]
		}
		return response, errors.New(response.Acknowledge())
	}

	return response, nil
}

func buildRequestURI(https bool, domain string, uri string) string {
	return fmt.Sprintf("%s://%s%s", httpSsl[https], domain, uri)
}
