package services

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/BNPrashanth/go-celeb-poc/internal/logger"
	"github.com/BNPrashanth/go-celeb-poc/models"
	"github.com/BNPrashanth/go-celeb-poc/restapi/operations/test"
)

// PaymentTest Func
func PaymentTest(params *test.GetTestParams) middleware.Responder {
	logger.Log.Info("Entering function..")
	return test.NewGetTestOK().WithPayload(&models.GeneralResponse{
		Success: true,
		Error:   nil,
		Message: "TEST method works..",
	})
}
