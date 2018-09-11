package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

var paypal = struct {
	APIKey string
}{
	APIKey: "sfsdf",
}

func init() {

	//Stripe API KEY
	stripe.Key = "sk_test_qv51kU2YKjb94IC8gdBPemu4"
}

//////////////////////////////////////
////////////Type Payment//////////////
//////////////////////////////////////

//Payment define a basic payment struct
type Payment struct {
	ID        string
	OrderID   int
	Method    string //webcheckout or CC data or tokenized
	Price     float32
	Currency  string
	Tax       float32
	ExtraData map[string]interface{}
	Token     string
}

//CreditCardPayment defines a struct to Credit Cards Charges
type CreditCardPayment struct {
	Payment
	CreditCard struct {
		CCV     string
		ExpDate string
		Number  string
	}
}

//WebCheckoutPayment defines a struct to WebChekcouts Payments
type WebCheckoutPayment struct {
	Payment
	Status string
}

////////////////////////////////////
//////CARDS WITHOUT REDIRECTS///////
////////////////////////////////////

//CreditCardStripe Generate a charge to a Credit Card
func (ccp *CreditCardPayment) CreditCardStripe() (err error) {

	//value on minimal unit, ej: "USD to cents"
	paymentAmount := ccp.Price * 100

	tax := ccp.Tax

	if tax > 0 {
		tax = (tax / 100) + 1
		paymentAmount += paymentAmount * tax
	}

	currency := strings.ToLower(ccp.Currency)

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(paymentAmount)),
		Currency:    stripe.String(currency),
		Description: stripe.String("Example charge"), //TODO: get services description
	}

	params.SetSource(ccp.Token)

	ch, err := charge.New(params)

	if stripeErr, ok := err.(*stripe.Error); ok {

		err = errors.New(stripeErr.Msg)
		return
	}

	ccp.ID = ch.ID

	return
}

////////////////////////////////////////////////
//////EXTERNAL CHECKOUT WITHOUT REDIRECTS///////
////////////////////////////////////////////////

func getPaypalBearerToken() (bearerToken string, err error) {
	req, err := http.NewRequest("GET", "https://api.sandbox.paypal.com/v1/oauth2/token", bytes.NewReader([]byte{}))

	if err != nil {
		return
	}

	req.Header.Add("client_id", paypal.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en_US")
	req.Header.Add("grant_type", "client_credentials")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return
	}

	responseData := struct {
		Scope       string `json:"scope"`
		Nonce       string `json:"nonce"`
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		AppID       string `json:"app_id"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	err = json.NewDecoder(res.Body).Decode(responseData)

	if err != nil {
		return
	}

	bearerToken = responseData.AccessToken

	return
}

//ButtonCheckoutPaypal execute a approved paypal pay
func (wcp *WebCheckoutPayment) ButtonCheckoutPaypal() (err error) {

	bearerToken, err := getPaypalBearerToken()

	if err != nil {
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/payments/payment/"+wcp.ID, bytes.NewReader([]byte{}))

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerToken)

	res, err := client.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {

		err = errors.New("Payment execution failed")
		return
	}

	var executedData map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(executedData)

	if err != nil {
		return
	}

	/* 	if executedData["intent"] != "sale" || executedData["state"] != "approved" {

		err = errors.New("Payment execution failed")
		return
	} */

	wcp.ID = executedData["id"].(string)

	return

}
