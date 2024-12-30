package database

import (
	"context"
	"fmt"

	"github.com/garupanojisan/tinylink/model"
	"github.com/garupanojisan/tinylink/pkg/tinylink/repository"
)

type tinyLinkImpl struct {
	mem map[int64]string
}

func NewTinyLinkRepository() repository.TinyLinkRepository {
	return &tinyLinkImpl{
		mem: make(map[int64]string),
	}
}

func (t *tinyLinkImpl) Create(ctx context.Context, m *model.TinyLink) error {
	t.mem[m.ID] = m.LongURL
	return nil
}

func (t *tinyLinkImpl) Find(ctx context.Context, id int64) (*model.TinyLink, error) {
	longURL, ok := t.mem[id]
	if !ok {
		return nil, fmt.Errorf("TinyLink not found for id %d", id)
	}
	return &model.TinyLink{ID: id, LongURL: longURL}, nil
}
