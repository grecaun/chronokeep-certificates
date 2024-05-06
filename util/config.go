package util

import (
	"image"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// GetConfig returns a config struct filled with values stored in local environment variables
func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Info(".env file wasn't loaded.")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil && port < 200 {
		port = 8181
	}

	development := os.Getenv("VERSION") != "production"

	file, err := os.Open(os.Getenv("CERTIFICATE_FILENAME"))
	if err != nil {
		return nil, errors.New("Unable to open file specified with CERTIFICATE_FILENAME environment variable.")
	}
	certificate, format, err := image.Decode(file)
	if err != nil {
		return nil, errors.New("Unable to decode file specified with CERTIFICATE_FILENAME environment variable.")
	}

	return &Config{
		Port:                   port,
		Development:            development,
		CertificateImage:       certificate,
		CertificateImageFormat: format,
	}, nil
}

// Config is the struct that holds all of the config values for connecting to a database
type Config struct {
	Port                   int
	Development            bool
	CertificateImage       image.Image
	CertificateImageFormat string
}
