package main

import (
	"github.com/BNPrashanth/go-celeb-poc/internal/config"
	"github.com/BNPrashanth/go-celeb-poc/internal/logger"
)

func main() {
	// Initialize Viper across the application
	config.InitializeViper()

	// Initialize Logger across the application
	logger.InitializeZapCustomLogger()

	logger.Log.Info("Starting Application..")

}
