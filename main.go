package main

import (
	"chronokeep/certificates/handlers"
	"chronokeep/certificates/util"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting certificate generator.")
	config, err := util.GetConfig()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	e := echo.New()
	e.Debug = config.Development

	log.Info("Setting up base middleware.")
	// Set up Recover and Logger middleware.
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  "${method} | ${status} | ${uri} -> ${latency_human}\n",
		Skipper: healthEndpointSkipper,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
	}))

	log.Info("Binding ")
	// Set up API handlers.
	handler := handlers.Handler{
		Config: config,
	}
	// Setup the Handler for validator
	handler.Bind(e.Group(""))

	e.Any("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	log.Info("Starting non https echo server.")
	log.Fatal(e.Start(":" + strconv.Itoa(config.Port)))
}

func healthEndpointSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/health")
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
