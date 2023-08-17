package fullurlsearch

import (
	"context"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
)

type FullUrlFind interface {
	FullUrlFind(context.Context, string) (linkentity.Link, error)
}

type FullUrl struct {
	storage FullUrlFind
}

func NewFullUrl(s FullUrlFind) *FullUrl {
	fu := &FullUrl{
		storage: s,
	}
	return fu
}

func (fu *FullUrl) FindByShort(ctx context.Context, shortUrl string) (linkentity.Link, error) {
	return fu.storage.FullUrlFind(ctx, shortUrl)
}
