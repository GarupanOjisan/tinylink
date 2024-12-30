package usecase

import (
	"github.com/labstack/echo/v4"

	"github.com/garupanojisan/tinylink/model"
	"github.com/garupanojisan/tinylink/pkg/tinylink/database"
	"github.com/garupanojisan/tinylink/pkg/tinylink/repository"
)

type TinyLink interface {
	CreateTinyLink(c echo.Context, longURL string) (string, error)
	GetLongURL(c echo.Context, encodedID string) (string, error)
}

type tinyLinkImpl struct {
	tinyLinkRepo repository.TinyLinkRepository
}

func NewTinyLink() TinyLink {
	return &tinyLinkImpl{
		tinyLinkRepo: database.NewTinyLinkRepository(),
	}
}

func (t *tinyLinkImpl) CreateTinyLink(c echo.Context, longURL string) (string, error) {
	m, err := model.NewTinyLink(longURL)
	if err != nil {
		return "", err
	}
	c.Logger().Info("CreateTinyLink", "id", m.ID, "longUrl", longURL)
	if err := t.tinyLinkRepo.Create(c.Request().Context(), m); err != nil {
		return "", err
	}
	return m.GetShortURL(), nil
}

func (t *tinyLinkImpl) GetLongURL(c echo.Context, base62EncodedID string) (string, error) {
	id := model.GetTinyLinkID(base62EncodedID)
	m, err := t.tinyLinkRepo.Find(c.Request().Context(), id)
	if err != nil {
		return "", err
	}
	return m.LongURL, nil
}
