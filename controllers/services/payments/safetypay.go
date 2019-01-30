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

type ExpressTokenRequest struct {
	XMLName             xml.Name `xml:"urn:CreateExpressToken"`
	APIKey              string   `xml:"urn:ApiKey,omitempty"`
	RequestDateTime     string   `xml:"urn:RequestDateTime,omitempty"`
	CurrencyID          string   `xml:"urn:CurrencyID,omitempty"`
	Amount              float32  `xml:"urn:Amount,omitempty"`
	MerchantSalesID     string   `xml:"urn:MerchantSalesID,omitempty"`
	Language            string   `xml:"urn:Language,omitempty"`
	ExpirationTime      string   `xml:"urn:ExpirationTime,omitempty"`
	FilterBy            string   `xml:"urn:FilterBy,omitempty"`
	TransactionOkURL    string   `xml:"urn:TransactionOkURL,omitempty"`
	TransactionErrorURL string   `xml:"urn:TransactionErrorURL,omitempty"`
	TrackingCode        string   `xml:"urn:TrackingCode,omitempty"`
	ProductID           string   `xml:"urn:ProductID,omitempty"`
	Signature           string   `xml:"urn:Signature,omitempty"`
}

type SafetyPayRequest struct {
	XMLName            xml.Name             `xml:"soapenv:Envelope"`
	SoapEnv            string               `xml:"xmlns:soapev,attr"`
	Urn                string               `xml:"xmlns:urn,attr"`
	CreateExpressToken *ExpressTokenRequest `xml:"soapenv:body>urn:CreateExpressToken"`
}

var safetyPay = struct {
	APIKey       string
	SignatureKey string
}{
	APIKey:       beego.AppConfig.String("safetypay::apiKey"),
	SignatureKey: beego.AppConfig.String("safetypay::signatureKey"),
}

func (s *SafetyPayRequest) createExpressTokenRequest() (token string, err error) {

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
	req.Header.Add("SOAPAction", "urn:safetypay:contract:mws:api:CommunicationTest")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New("Error with token request")
		//return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))
	return

}

func SafetyPayCreateExpressToken(currencyID string, amount float32, orderID int, transactionOkURL string, transactionErrorURL string, filter string) (expressToken string, err error) {

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

	fmt.Println(amountString)

	signature := sha256.Sum256([]byte(requestDateTime + currencyID + amountString + strconv.Itoa(orderID) + "ES" + "X" + "120" + transactionOkURL + transactionErrorURL + safetyPay.SignatureKey))

	signatureStr := string(signature[:32])

	tokenStruct := &ExpressTokenRequest{
		APIKey:              safetyPay.APIKey,
		RequestDateTime:     requestDateTime,
		CurrencyID:          currencyID,
		Amount:              amount,
		MerchantSalesID:     strconv.Itoa(orderID),
		FilterBy:            filterBy,
		ProductID:           productID,
		TransactionOkURL:    transactionOkURL,
		TransactionErrorURL: transactionErrorURL,
		TrackingCode:        "X",
		ExpirationTime:      "120",
		Signature:           signatureStr,
	}

	safetyPayStruct := &SafetyPayRequest{
		SoapEnv:            "http://schemas.xmlsoap.org/soap/envelope/",
		Urn:                "urn:safetypay:messages:mws:api",
		CreateExpressToken: tokenStruct,
	}

	expressToken, err = safetyPayStruct.createExpressTokenRequest()

	return

}

//type SafetyPayResponse struct
