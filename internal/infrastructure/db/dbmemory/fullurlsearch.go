package dbmemory

import (
	"context"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/search/fullurlsearch"
)

var _ fullurlsearch.FullUrlFind = &FullUrlFind{}

type FullUrlFind struct{}

func NewFullUrlFind() *FullUrlFind {
	return &FullUrlFind{}
}

func (db *FullUrlFind) FullUrlFind(ctx context.Context, shortUrl string) (linkentity.Link, error) {
	return linkentity.Link{}, nil
}
