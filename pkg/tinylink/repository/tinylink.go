package repository

import (
	"context"
	"github.com/garupanojisan/tinylink/model"
)

type TinyLinkRepository interface {
	Create(ctx context.Context, m *model.TinyLink) error
	Find(ctx context.Context, id int64) (*model.TinyLink, error)
}
