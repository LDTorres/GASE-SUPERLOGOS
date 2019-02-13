package payments

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego"
)

var (
	APIKey       = beego.AppConfig.String("safetypay::apiKey")
	SignatureKey = beego.AppConfig.String("safetypay::signatureKey")
)

// SafetyPayRequest describe a SafetyPay xml Env
type SafetyPayRequest struct {
	XMLName                     xml.Name                                     `xml:"soap:Envelope"`
	SoapEnv                     string                                       `xml:"xmlns:soap,attr"`
	Urn                         string                                       `xml:"xmlns:urn,attr"`
	Urn1                        string                                       `xml:"xmlns:urn1,attr"`
	Header                      string                                       `xml:"soap:Header"`
	CreateExpressToken          *SafetyPayExpressTokenRequest                `xml:"soap:Body>urn:ExpressTokenRequest,omitempty"`
	OperationActivity           *SafetyPayOperationActivityRequest           `xml:"soap:Body>urn:OperationActivityRequest,omitempty"`
	ConfirmNewOperationActivity *SafetyPayConfirmNewOperationActivityRequest `xml:"soap:Body>urn:OperationActivityNotifiedRequest,omitempty"`
}

// SafetyPayResponse describe a SafetyPay xml Env
type SafetyPayResponse struct {
	XMLName                     xml.Name                                      `xml:"Envelope"`
	CreateExpressToken          *SafetyPayExpressTokenResponse                `xml:"Body>ExpressTokenResponse,omitempty"`
	OperationActivity           *SafetyPayOperationActivityResponse           `xml:"Body>OperationResponse>ListOfOperations,omitempty"`
	ConfirmNewOperationActivity *SafetyPayConfirmNewOperationActivityResponse `xml:"Body>OperationActivityNotifiedResponse,omitempty"`
}

// SafetyPayExpressTokenRequest describe a SafetyPay Express Token Request
type SafetyPayExpressTokenRequest struct {
	XMLName                   xml.Name `xml:"urn:ExpressTokenRequest"`
	APIKey                    string   `xml:"urn:ApiKey,omitempty"`
	RequestDateTime           string   `xml:"urn:RequestDateTime,omitempty"`
	CurrencyID                string   `xml:"urn:CurrencyID,omitempty"`
	Amount                    string   `xml:"urn:Amount,omitempty"`
	MerchantSalesID           string   `xml:"urn:MerchantSalesID,omitempty"`
	Language                  string   `xml:"urn1:Language,omitempty"`
	TrackingCode              string   `xml:"urn:TrackingCode"`
	ExpirationTime            string   `xml:"urn:ExpirationTime,omitempty"`
	TransactionExpirationTime string   `xml:"urn:TransactionExpirationTime,omitempty"`
	FilterBy                  string   `xml:"urn:FilterBy,omitempty"`
	TransactionOkURL          string   `xml:"urn:TransactionOkURL,omitempty"`
	TransactionErrorURL       string   `xml:"urn:TransactionErrorURL,omitempty"`
	ProductID                 string   `xml:"urn:ProductID,omitempty"`
	Signature                 string   `xml:"urn:Signature,omitempty"`
}

// SafetyPayExpressTokenResponse ...
type SafetyPayExpressTokenResponse struct {
	XMLName            xml.Name `xml:"ExpressTokenResponse"`
	ResponseDateTime   string   `xml:"ResponseDateTime,omitempty"`
	ShopperRedirectURL string   `xml:"ShopperRedirectURL,omitempty"`
	Signature          string   `xml:"Signature,omitempty"`
	ErrorManager       string   `xml:"ErrorManager>Description,omitempty"`
}

// SafetyPayOperationActivityRequest describe a Safetypay operation activity request
type SafetyPayOperationActivityRequest struct {
	XMLName         xml.Name `xml:"urn:OperationActivityRequest"`
	APIKey          string   `xml:"urn:ApiKey,omitempty"`
	RequestDateTime string   `xml:"urn:RequestDateTime,omitempty"`
	Signature       string   `xml:"urn:Signature,omitempty"`
}

// SafetyPayOperationActivityResponse ...
type SafetyPayOperationActivityResponse struct {
	XMLName    xml.Name                                       `xml:"ListOfOperations"`
	Operations []*SafetyPayOperationActivityResponseOperation `xml:"Operation"`
}

// SafetyPayOperationActivityResponseOperation ...
type SafetyPayOperationActivityResponseOperation struct {
	XMLName          xml.Name `xml:"Operation"`
	CreationDateTime string   `xml:"CreationDateTime,omitempty"`
	OperationID      string   `xml:"OperationID,omitempty"`
	MerchantSalesID  string   `xml:"MerchantSalesID,omitempty"`
	Status           string   `xml:"OperationActivities>OperationActivity>Status>StatusCode,omitempty"`
}

// SafetyPayOperationActivity describe a operation Object from Safety Pay
type SafetyPayOperationActivity struct {
	CreationDateTime   string  `xml:"urn:CreationDateTime"`
	OperationID        string  `xml:"urn:OperationID"`
	MerchantSalesID    string  `xml:"urn:MerchantSalesID"`
	MerchantOrderID    string  `xml:"urn:MerchantOrderID"`
	Amount             float32 `xml:"urn:Amount"`
	CurrencyID         string  `xml:"urn:CurrencyID"`
	ShopperAmount      float32 `xml:"urn:ShopperAmount"`
	ShopperCurrencyID  string  `xml:"urn:ShopperCurrencyID"`
	AuthorizationCode  string  `xml:"urn:AuthorizationCode,omitempty"`
	PaymentReferenceNo string  `xml:"urn:PaymentReferenceNo,omitempty"`
	OperationStatus    string  `xml:"urn:OperationStatus,omitempty"`
}

type ConfirmOperation struct {
	Operations []*SafetyPayConfirmOperationActivity `xml:"urn1:ConfirmOperation"`
}

// SafetyPayConfirmNewOperationActivityRequest confirm multiple operations
type SafetyPayConfirmNewOperationActivityRequest struct {
	XMLName             xml.Name          `xml:"urn:OperationActivityNotifiedRequest"`
	APIKey              string            `xml:"urn:ApiKey,omitempty"`
	RequestDateTime     string            `xml:"urn:RequestDateTime,omitempty"`
	OperationActivities *ConfirmOperation `xml:"urn:ListOfOperationsActivityNotified,omitempty"`
	Signature           string            `xml:"urn:Signature,omitempty"`
}

// SafetyPayConfirmNewOperationActivityResponse ...
type SafetyPayConfirmNewOperationActivityResponse struct {
	XMLName          xml.Name `xml:"OperationActivityNotifiedResponse"`
	ResponseDateTime string   `xml:"ResponseDateTime,omitempty"`
	Signature        string   `xml:"Signature,omitempty"`
	ErrorManager     string   `xml:"ErrorManager>Description,omitempty"`
}

// SafetyPayConfirmOperationActivity describe a confirm operation Object from Safety Pay
type SafetyPayConfirmOperationActivity struct {
	CreationDateTime string `xml:"urn1:CreationDateTime"`
	OperationID      string `xml:"urn1:OperationID"`
	MerchantSalesID  string `xml:"urn1:MerchantSalesID"`
	MerchantOrderID  string `xml:"urn1:MerchantOrderID"`
	OperationStatus  string `xml:"urn1:OperationStatus,omitempty"`
}

func createSignature256(args ...string) (signature256 string) {

	var stringArgs string

	for _, arg := range args {
		stringArgs += arg
	}

	input := strings.NewReader(stringArgs)

	hash := sha256.New()
	if _, err := io.Copy(hash, input); err != nil {
		log.Fatal(err)
	}

	sha := hex.EncodeToString(hash.Sum(nil))

	// beego.Debug(stringArgs)
	// beego.Debug(sha)

	signature256 = sha
	return
}

func (s *SafetyPayRequest) createExpressTokenRequest() (URL *SafetyPayResponse, err error) {

	output, err := xml.MarshalIndent(s, "  ", "    ")
	if err != nil {
		return
	}

	requestBodyData := bytes.NewReader(output)

	// fmt.Println(string(output))

	req, err := http.NewRequest("POST", "https://sandbox-mws2.safetypay.com/express/ws/v.3.0/", requestBodyData)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "urn:safetypay:contract:mws:api:CreateExpressToken")

	beego.Debug("Req Headers: ", req.Header.Get("SOAPAction"))

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		beego.Debug("ERR: ", err.Error())
		return
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		beego.Debug("ERR  2: ", err.Error())
		return
	}

	r := &SafetyPayResponse{}

	err = xml.Unmarshal([]byte(data), &r)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(&r)

	if res.StatusCode != 200 {
		beego.Debug("Response: ", res.Body)
		err = errors.New("Error with token request")
		return
	}

	return r, nil
}

// SafetyPayCreateExpressToken ...
func SafetyPayCreateExpressToken(currencyID string, amount float32, orderID int, transactionOkURL string, transactionErrorURL string, filter string) (expressToken *SafetyPayResponse, err error) {

	//condicional filter online o efectivo
	var (
		filterBy  string
		productID string
	)

	if filter == "online" {
		filterBy = "CHANNEL(OL)"
		productID = "1"
	} else if filter == "cash" {
		filterBy = "CHANNEL(WP)"
		productID = "2"
	} else {
		err = errors.New("filter value is not valid")
		return
	}

	requestDateTime := jodaTime.Format("yyyy-MM-ddThh:mm:ss", time.Now())
	amountString := strconv.FormatFloat(float64(amount), 'f', 2, 32)

	signature := createSignature256(requestDateTime, currencyID, amountString, strconv.Itoa(orderID), "ES", "", "120", transactionOkURL, transactionErrorURL, SignatureKey)

	// beego.Debug("requestDateTime: ", requestDateTime)
	// beego.Debug("Signature: ", signature)
	// beego.Debug("APIKey: ", APIKey)

	tokenStruct := &SafetyPayExpressTokenRequest{
		APIKey:                    APIKey,
		RequestDateTime:           requestDateTime,
		CurrencyID:                "EUR",
		Amount:                    amountString,
		MerchantSalesID:           strconv.Itoa(orderID),
		Language:                  "ES",
		TrackingCode:              "",
		ExpirationTime:            "120",
		FilterBy:                  filterBy,
		TransactionOkURL:          transactionOkURL,
		TransactionErrorURL:       transactionErrorURL,
		TransactionExpirationTime: "120",
		ProductID:                 productID,
		Signature:                 signature,
	}

	// beego.Debug("tokenStruct: ", tokenStruct)

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv:            "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:                "urn:safetypay:messages:mws:api",
		Urn1:               "urn:safetypay:schema:mws:api",
		CreateExpressToken: tokenStruct,
	}

	// beego.Debug("safetyPayStruct: ", safetyPayStruct)

	expressToken, err = safetyPayStruct.createExpressTokenRequest()

	if err != nil {
		beego.Debug("createExpressTokenRequest ERR ", err)
		return
	}

	beego.Debug("expressToken: ", expressToken)

	return
}

// getNewOperationActivityRequest GET NEW PAID ORDERS
func (s *SafetyPayRequest) getNewOperationActivityRequest() (r *SafetyPayResponse, err error) {

	output, err := xml.MarshalIndent(s, "  ", "    ")
	if err != nil {
		return
	}

	requestBodyData := bytes.NewReader(output)

	req, err := http.NewRequest("POST", "https://sandbox-mws2.safetypay.com/express/ws/v.3.0/", requestBodyData)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "urn:safetypay:contract:mws:api:GetNewOperationActivity")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New("Error with token request")
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	r = &SafetyPayResponse{}

	err = xml.Unmarshal([]byte(body), &r)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(&r)

	if res.StatusCode != 200 {
		beego.Debug("Response: ", res.Body)
		err = errors.New("Error with token request")
		return
	}

	return r, nil
}

// SafetyPayGetNewOperationActivity ...
func SafetyPayGetNewOperationActivity() (response *SafetyPayResponse, err error) {

	requestDateTime := jodaTime.Format("yyyy-MM-ddThh:mm:ss", time.Now())

	signature := createSignature256(requestDateTime, SignatureKey)

	operationStruct := &SafetyPayOperationActivityRequest{
		APIKey:          APIKey,
		RequestDateTime: requestDateTime,
		Signature:       signature,
	}

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv:           "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:               "urn:safetypay:messages:mws:api",
		Urn1:              "urn:safetypay:schema:mws:api",
		OperationActivity: operationStruct,
	}

	response, err = safetyPayStruct.getNewOperationActivityRequest()

	return
}

// confirmNewOperationActivity CONFIRM NEW PAID ORDERS
func (s *SafetyPayRequest) confirmNewOperationActivityRequest() (r *SafetyPayResponse, err error) {

	output, err := xml.MarshalIndent(s, "  ", "    ")
	if err != nil {
		return
	}

	fmt.Println(string(output))

	requestBodyData := bytes.NewReader(output)

	req, err := http.NewRequest("POST", "https://sandbox-mws2.safetypay.com/express/ws/v.3.0/", requestBodyData)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "urn:safetypay:contract:mws:api:ConfirmNewOperationActivity")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New("Error with token request")
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	r = &SafetyPayResponse{}

	err = xml.Unmarshal([]byte(body), &r)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(&r)

	if res.StatusCode != 200 {
		beego.Debug("Response: ", res.Body)
		err = errors.New("Error with token request")
		return
	}

	return r, nil
}

// SafetypayConfirmNewOperationActivity ...
func SafetypayConfirmNewOperationActivity(operationActivities []*SafetyPayConfirmOperationActivity) (r *SafetyPayResponse, err error) {

	requestDateTime := jodaTime.Format("yyyy-MM-ddThh:mm:ss", time.Now())

	operationActivitiesStrings := []string{requestDateTime}

	for _, operationActivity := range operationActivities {
		operationActivitiesStrings = append(operationActivitiesStrings, operationActivity.OperationID, operationActivity.MerchantSalesID, operationActivity.MerchantOrderID, operationActivity.OperationStatus)
	}

	operationActivitiesStrings = append(operationActivitiesStrings, SignatureKey)

	signature := createSignature256(operationActivitiesStrings...)

	confirmOperations := &ConfirmOperation{
		Operations: operationActivities,
	}

	operationStruct := &SafetyPayConfirmNewOperationActivityRequest{
		APIKey:              APIKey,
		RequestDateTime:     requestDateTime,
		Signature:           signature,
		OperationActivities: confirmOperations,
	}

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:     "urn:safetypay:messages:mws:api",
		Urn1:    "urn:safetypay:schema:mws:api",
		ConfirmNewOperationActivity: operationStruct,
	}

	r, err = safetyPayStruct.confirmNewOperationActivityRequest()

	return
}
