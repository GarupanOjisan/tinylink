package model

import (
	"github.com/garupanojisan/tinylink/util"
	"os"
)

type TinyLink struct {
	ID      int64
	LongURL string
}

func NewTinyLink(longURL string) (*TinyLink, error) {
	id, err := util.GenerateUniqueID()
	if err != nil {
		return nil, err
	}

	m := &TinyLink{
		ID:      id,
		LongURL: longURL,
	}
	return m, nil
}

func (m *TinyLink) GetBase62EncodedID() string {
	return util.EncodeBase62(m.ID)
}

func (m *TinyLink) GetShortURL() string {
	if os.Getenv("SITE_DOMAIN") == "" {
		panic("SITE_DOMAIN not set")
	}
	return os.Getenv("SITE_DOMAIN") + "/" + m.GetBase62EncodedID()
}

func GetTinyLinkID(base62EncodedID string) int64 {
	return util.DecodeBase62(base62EncodedID)
}
