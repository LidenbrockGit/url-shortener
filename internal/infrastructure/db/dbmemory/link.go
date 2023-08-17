package dbmemory

import (
	"context"
	"fmt"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/linkrepo"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/search/fullurlsearch"
	"github.com/google/uuid"
)

var _ linkrepo.Storage = &LinkRepo{}
var _ fullurlsearch.FullUrlFind = &LinkRepo{}

type LinkRepo struct {
	memStorage map[uuid.UUID]linkentity.Link
}

func NewLinkRepo() *LinkRepo {
	return &LinkRepo{
		memStorage: make(map[uuid.UUID]linkentity.Link),
	}
}

func (db *LinkRepo) FullUrlFind(ctx context.Context, shortUrl string) (linkentity.Link, error) {
	select {
	case <-ctx.Done():
		return linkentity.Link{}, ctx.Err()
	default:
	}

	var resLink linkentity.Link
	for _, link := range db.memStorage {
		if !link.IsDeleted() && link.ShortUrl == shortUrl {
			resLink = link
			break
		}
	}

	if resLink == (linkentity.Link{}) {
		return linkentity.Link{}, LinkFindErr
	}

	resLink.UsageCount++
	err := db.Update(ctx, resLink)
	if err != nil {
		return linkentity.Link{}, LinkUpdateErr
	}

	return resLink, nil
}

func (db *LinkRepo) Create(ctx context.Context, link linkentity.Link) (linkentity.Link, error) {
	select {
	case <-ctx.Done():
		return linkentity.Link{}, ctx.Err()
	default:
	}

	for _, l := range db.memStorage {
		if !l.IsDeleted() && l.ShortUrl == link.ShortUrl {
			return linkentity.Link{}, fmt.Errorf("%w; shortUrl '%s'", LinkShortUrlDuplicatingErr, link.ShortUrl)
		}
	}

	link.Id = uuid.New()
	link.CreatedAt = time.Now()
	db.memStorage[link.Id] = link

	return link, nil
}

func (db *LinkRepo) Read(ctx context.Context, linkId uuid.UUID) (linkentity.Link, error) {
	select {
	case <-ctx.Done():
		return linkentity.Link{}, ctx.Err()
	default:
	}

	link, ok := db.memStorage[linkId]
	if !ok || link.IsDeleted() {
		return linkentity.Link{}, fmt.Errorf("%w; linkId '%s'", LinkFindErr, linkId)
	}

	return link, nil
}

func (db *LinkRepo) ReadAll(ctx context.Context) (<-chan linkentity.Link, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	chOut := make(chan linkentity.Link)
	go func() {
		for _, link := range db.memStorage {
			if link.IsDeleted() {
				continue
			}
			chOut <- link
		}
		close(chOut)
	}()

	return chOut, nil
}

func (db *LinkRepo) Update(ctx context.Context, link linkentity.Link) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := db.Read(ctx, link.Id)
	if err != nil {
		return err
	}

	db.memStorage[link.Id] = link

	return nil
}

func (db *LinkRepo) Delete(ctx context.Context, linkId uuid.UUID) (linkentity.Link, error) {
	select {
	case <-ctx.Done():
		return linkentity.Link{}, ctx.Err()
	default:
	}

	deleteLink, err := db.Read(ctx, linkId)
	if err != nil {
		return linkentity.Link{}, err
	}

	deleteLink.DeletedAt = time.Now()
	err = db.Update(ctx, deleteLink)
	if err != nil {
		return linkentity.Link{}, err
	}

	return deleteLink, nil
}
