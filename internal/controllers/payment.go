package controllers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/BNPrashanth/go-celeb-poc/internal/logger"
	"github.com/BNPrashanth/go-celeb-poc/models"
	"github.com/BNPrashanth/go-celeb-poc/restapi/operations/test"
)

// HandleTest Func
func HandleTest(params *test.GetTestParams) middleware.Responder {
	logger.Log.Info("TEST method called..")
	return test.NewGetTestOK().WithPayload(&models.GeneralResponse{
		Success: true,
		Error:   nil,
		Message: "TEST method works..",
	})
}
