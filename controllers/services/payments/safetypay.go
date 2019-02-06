package payments

import (
	"bytes"
	"crypto/sha256"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/vjeantet/jodaTime"

	"github.com/astaxie/beego"
)

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

//TODO: AGFREGAR ENMARCADO XML BASADO EN Manual

// SafetyPayOperationActivityRequest describe a Safetypay operation activity request
type SafetyPayOperationActivityRequest struct {
	XMLName         xml.Name `xml:"urn:GetNewOperationActivity"`
	APIKey          string   `xml:"urn:ApiKey,omitempty"`
	RequestDateTime string   `xml:"urn:RequestDateTime,omitempty"`
	Signature       string   `xml:"urn:Signature,omitempty"`
}

// SafetyPayConfirmNewOperationActivityRequest confirm multiple operations
type SafetyPayConfirmNewOperationActivityRequest struct {
	XMLName             xml.Name                      `xml:"urn:GetNewOperationActivity"`
	APIKey              string                        `xml:"urn:ApiKey,omitempty"`
	RequestDateTime     string                        `xml:"urn:RequestDateTime,omitempty"`
	OperationActivities []*SafetyPayOperationActivity `xml:"urn:ListOfOperationsActivityNotified,omitempty"`
	Signature           string                        `xml:"urn:Signature,omitempty"`
}

// SafetyPayExpressTokenRequest describe a SafetyPay Express Token Request
type SafetyPayExpressTokenRequest struct {
	XMLName                   xml.Name `xml:"urn:CreateExpressToken"`
	APIKey                    string   `xml:"urn:ApiKey,omitempty"`
	RequestDateTime           string   `xml:"urn:RequestDateTime,omitempty"`
	CurrencyID                string   `xml:"urn:CurrencyID,omitempty"`
	Amount                    float32  `xml:"urn:Amount,omitempty"`
	MerchantSalesID           string   `xml:"urn:MerchantSalesID,omitempty"`
	Language                  string   `xml:"urn1:Language,omitempty"`
	ExpirationTime            string   `xml:"urn:ExpirationTime,omitempty"`
	TransactionExpirationTime string   `xml:"urn:TransactionExpirationTime,omitempty"`
	FilterBy                  string   `xml:"urn:FilterBy,omitempty"`
	TransactionOkURL          string   `xml:"urn:TransactionOkURL,omitempty"`
	TransactionErrorURL       string   `xml:"urn:TransactionErrorURL,omitempty"`
	TrackingCode              string   `xml:"urn:TrackingCode,omitempty"`
	ProductID                 string   `xml:"urn:ProductID,omitempty"`
	Signature                 string   `xml:"urn:Signature,omitempty"`
}

// SafetyPayRequest describe a SafetyPay xml Env
type SafetyPayRequest struct {
	XMLName                     xml.Name `xml:"soapenv:Envelope"`
	SoapEnv                     string   `xml:"xmlns:soapev,attr"`
	Urn                         string   `xml:"xmlns:urn,attr"`
	Urn1                        string   `xml:"xmlns:urn1,attr,omitempty"`
	SOAPAction                  string
	CreateExpressToken          *SafetyPayExpressTokenRequest                `xml:"soapenv:Body>urn:ExpressTokenRequest"`
	OperationActivity           *SafetyPayOperationActivityRequest           `xml:"soapenv:Body>urn:OperationActivityRequest"`
	ConfirmNewOperationActivity *SafetyPayConfirmNewOperationActivityRequest `xml:"soapenv:Body>urn:OperationActivityNotifiedRequest"`
}

var safetyPay = struct {
	APIKey       string
	SignatureKey string
}{
	APIKey:       beego.AppConfig.String("safetypay::apiKey"),
	SignatureKey: beego.AppConfig.String("safetypay::signatureKey"),
}

func createSignature256(args ...string) (signature256 string) {

	var stringArgs string

	for _, arg := range args {
		stringArgs += arg
	}

	argsBytes := []byte(stringArgs)

	signature := sha256.Sum256(argsBytes)

	signature256 = string(signature[:32]) //DEVUELVE CODIFICACION RARA

	return
}

func (s *SafetyPayRequest) createExpressTokenRequest() (URL string, err error) {

	output, err := xml.MarshalIndent(s, "  ", "    ")
	if err != nil {
		return
	}

	requestBodyData := bytes.NewReader(output)

	fmt.Println(string(output))

	req, err := http.NewRequest("POST", "https://sandbox-mws2.safetypay.com/express/ws/v.3.0/", requestBodyData)

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("SOAPAction", "urn:safetypay:contract:mws:api:"+s.SOAPAction)

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
	URLBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	URL = string(URLBytes)

	fmt.Println(URL)

	/**** OBTENER TOKEN/URL DEL CUERPO DE LA PETICION ****/

	return

}

// SafetyPayCreateExpressToken ...
func SafetyPayCreateExpressToken(currencyID string, amount float32, orderID int, transactionOkURL string, transactionErrorURL string, filter string) (expressToken string, err error) {

	//condicional filter online o efectivo
	var (
		// filterBy  string
		productID string
	)

	if filter == "online" {
		// filterBy = "CHANNEL(OL)"
		productID = "1"
	} else if filter == "cash" {
		// filterBy = "CHANNEL(WP)"
		productID = "2"
	} else {
		err = errors.New("filter value is not valid")
		return
	}

	requestDateTime := jodaTime.Format("yyyy-MM-ddThh:mm:ss", time.Now())
	amountString := strconv.FormatFloat(float64(amount), 'f', 2, 32)

	signature := createSignature256(requestDateTime, currencyID, amountString, strconv.Itoa(orderID), "ES", "", "120", transactionOkURL, transactionErrorURL, safetyPay.SignatureKey)

	// TODO: REVISAR SIGNATURE RESULTANTE, el string posee una codificacion no valida

	tokenStruct := &SafetyPayExpressTokenRequest{
		APIKey:          safetyPay.APIKey,
		RequestDateTime: requestDateTime,
		CurrencyID:      "EUR",
		Amount:          amount,
		MerchantSalesID: strconv.Itoa(orderID),
		Language:        "ES",
		TrackingCode:    "",
		ExpirationTime:  "120",
		// FilterBy:                  filterBy,
		TransactionOkURL:          transactionOkURL,
		TransactionErrorURL:       transactionErrorURL,
		TransactionExpirationTime: "120",
		ProductID:                 productID,
		Signature:                 signature,
	}

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv:            "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:                "urn:safetypay:messages:mws:api",
		Urn1:               "urn:safetypay:schema:mws:api",
		CreateExpressToken: tokenStruct,
		SOAPAction:         "CreateExpressToken",
	}

	expressToken, err = safetyPayStruct.createExpressTokenRequest()

	return

}

func (s *SafetyPayRequest) getNewOperationActivity() (operationActivities []*SafetyPayOperationActivity, err error) {

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
	req.Header.Add("SOAPAction", "urn:safetypay:contract:mws:api:"+s.SOAPAction)

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

	fmt.Println(body)

	//TODO: RECUPERAR DEL CUERPO DE LA RESPUESTA las operationsActivities

	return
}

// SafetyPayGetNewOperationActivity ...
func SafetyPayGetNewOperationActivity() (operationActivities []*SafetyPayOperationActivity, err error) {

	requestDateTime := jodaTime.Format("yyyy-MM-ddThh:mm:ss", time.Now())

	signature := createSignature256(requestDateTime, safetyPay.SignatureKey)

	operationStruct := &SafetyPayOperationActivityRequest{
		APIKey:          safetyPay.APIKey,
		RequestDateTime: requestDateTime,
		Signature:       signature,
	}

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv:           "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:               "urn:safetypay:messages:mws:api",
		OperationActivity: operationStruct,
		SOAPAction:        "GetNewOperationActivity",
	}

	operationActivities, err = safetyPayStruct.getNewOperationActivity()

	return

}

// SafetypayConfirmNewOperationActivity ...
func SafetypayConfirmNewOperationActivity(operationActivities []*SafetyPayOperationActivity) (err error) {

	requestDateTime := jodaTime.Format("yyyy-MM-ddThh:mm:ss", time.Now())

	operationActivitiesStrings := []string{requestDateTime}

	for _, operationActivity := range operationActivities {
		operationActivitiesStrings = append(operationActivitiesStrings, operationActivity.OperationID, operationActivity.MerchantSalesID, operationActivity.MerchantOrderID, operationActivity.OperationStatus)
	}

	operationActivitiesStrings = append(operationActivitiesStrings, safetyPay.SignatureKey)

	signature := createSignature256(operationActivitiesStrings...)

	operationStruct := &SafetyPayConfirmNewOperationActivityRequest{
		APIKey:          safetyPay.APIKey,
		RequestDateTime: requestDateTime,
		Signature:       signature,
	}

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:     "urn:safetypay:messages:mws:api",
		ConfirmNewOperationActivity: operationStruct,
		SOAPAction:                  "ConfirmNewOperationActivity",
	}

	operationActivities, err = safetyPayStruct.getNewOperationActivity()

	return

}
