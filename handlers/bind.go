package handlers

import (
	"chronokeep/certificates/util"

	"github.com/labstack/echo/v4"
)

// Handler Struct for using methods for handling information.
type Handler struct {
	Config *util.Config
}

func (h Handler) Bind(group *echo.Group) {
	// Certificate image
	group.GET("/:name/:event/:time/:date", h.GetCertificate)
}
