package payments

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
)

//PAYPAL API DATA
var paypal = struct {
	APIKey string
	Secret string
}{
	APIKey: beego.AppConfig.String("paypal::apiKey"),
	Secret: beego.AppConfig.String("paypal::secret"),
}

type paypalResponseToken struct {
	Scope       string `json:"scope"`
	Nonce       string `json:"nonce"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppID       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
}

////////////////////////////////////////////////
//////EXTERNAL CHECKOUT WITHOUT REDIRECTS///////
////////////////////////////////////////////////

func getPaypalBearerToken() (bearerToken string, err error) {

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	requestBodyData := data.Encode()

	req, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/oauth2/token", strings.NewReader(requestBodyData))

	if err != nil {
		return
	}

	req.SetBasicAuth(paypal.APIKey, paypal.Secret)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en_US")
	req.Header.Add("grant_type", "client_credentials")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = errors.New("Error with token request")
		return
	}

	responseData := &paypalResponseToken{}

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

	requestBodyData := &struct {
		PayerID string `json:"payer_id"`
	}{PayerID: wcp.PayerID}

	requestBodyBytes, _ := json.Marshal(requestBodyData)

	req, err := http.NewRequest("POST", "https://api.sandbox.paypal.com/v1/payments/payment/"+wcp.ID+"/execute", bytes.NewReader(requestBodyBytes))

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

	executedData := map[string]interface{}{}

	err = json.NewDecoder(res.Body).Decode(&executedData)

	if err != nil {
		return
	}

	/* 	if executedData["intent"] != "sale" || executedData["state"] != "approved" {

		err = errors.New("Payment execution failed")
		return
	} */
	//TODO: arreglar index
	//wcp.ID = executedData["id"].(string)

	return

}
