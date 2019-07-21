package services

import (
	"bytes"
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/BNPrashanth/go-celeb-poc/internal/logger"
	"github.com/BNPrashanth/go-celeb-poc/models"
	"github.com/BNPrashanth/go-celeb-poc/restapi/operations/test"
)

// PPClient struct
type PPClient struct {
	Scope       string
	AccessToken string `json:"access_token"`
	Type        string
	AppID       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

// PayPalClient VAR
var PayPalClient *PPClient

// IntitalizePayPalService Func
func IntitalizePayPalService() {
	// Create a paypal client instance
	req, err := http.NewRequest(
		"POST",
		"https://api.sandbox.paypal.com/v1/oauth2/token",
		bytes.NewBuffer([]byte("grant_type=client_credentials")),
	)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	err = requestWithBasic(req, PayPalClient)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	if PayPalClient.AccessToken != "" {
		logger.Log.Info("PayPal Service started..")
	}

}

// PaymentTest Func
func PaymentTest(params *test.GetTestParams) middleware.Responder {
	logger.Log.Info("Entering function..")

	return test.NewGetTestOK().WithPayload(&models.GeneralResponse{
		Success: true,
		Error:   nil,
		Message: "TEST method works.. Access Token: " + PayPalClient.AccessToken,
	})
}
