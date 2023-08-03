package linkrepo

import (
	"context"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/google/uuid"
)

const readChanSize = 512
const readLimitWaitSec = 2

type Storage interface {
	Create(ctx context.Context, link linkentity.Link) (uuid.UUID, error)
	Read(ctx context.Context, linkId uuid.UUID) (linkentity.Link, error)
	ReadAll(ctx context.Context) (<-chan linkentity.Link, error)
	Update(ctx context.Context, link linkentity.Link) error
	Delete(ctx context.Context, linkId uuid.UUID) (linkentity.Link, error)
}

type LinkRepo struct {
	storage Storage
}

func NewLinkRepo(s Storage) LinkRepo {
	l := LinkRepo{
		storage: s,
	}
	return l
}

func (l *LinkRepo) Create(ctx context.Context, link linkentity.Link) (uuid.UUID, error) {
	return l.storage.Create(ctx, link)
}

func (l *LinkRepo) Read(ctx context.Context, id uuid.UUID) (linkentity.Link, error) {
	return l.storage.Read(ctx, id)
}

func (l *LinkRepo) ReadAll(ctx context.Context) (<-chan linkentity.Link, error) {
	chOut := make(chan linkentity.Link, readChanSize)

	chin, err := l.storage.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(chOut)
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second * readLimitWaitSec):
				return
			case link, ok := <-chin:
				if !ok {
					return
				}
				chOut <- link
			}
		}
	}()

	return chOut, nil
}

func (l *LinkRepo) Update(ctx context.Context, link linkentity.Link) error {
	return l.storage.Update(ctx, link)
}

func (l *LinkRepo) Delete(ctx context.Context, id uuid.UUID) (linkentity.Link, error) {
	return l.storage.Delete(ctx, id)
}
