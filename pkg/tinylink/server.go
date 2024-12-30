package tinylink

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/garupanojisan/tinylink/pkg/tinylink/usecase"
)

type Server struct {
	useCase usecase.TinyLink
}

func NewTinyLinkServer() *Server {
	return &Server{
		useCase: usecase.NewTinyLink(),
	}
}

func (s *Server) RedirectToLongURL(c echo.Context) error {
	encodedID := c.Param("id")
	longURL, err := s.useCase.GetLongURL(c, encodedID)
	if err != nil {
		return err
	}
	return c.Redirect(301, longURL)
}

type (
	CreateTinyLinkRequest struct {
		URL string `json:"url" validate:"required,url"`
	}
	CreateTinyLinkResponse struct {
		ShortURL string `json:"short_url"`
	}
)

func (s *Server) CreateNewTinyLink(c echo.Context) error {
	req := new(CreateTinyLinkRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	shortURL, err := s.useCase.CreateTinyLink(c, req.URL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, CreateTinyLinkResponse{
		ShortURL: shortURL,
	})
}
