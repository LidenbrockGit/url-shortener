package userrepo

import (
	"context"

	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/google/uuid"
)

type Storage interface {
	Create(ctx context.Context, user userentity.User) (uuid.UUID, error)
	Read(ctx context.Context, userId uuid.UUID) (userentity.User, error)
	Update(ctx context.Context, user userentity.User) error
	Delete(ctx context.Context, userId uuid.UUID) (userentity.User, error)
	Search(ctx context.Context, login string) (userentity.User, error)
}

type UserRepo struct {
	storage Storage
}

func NewUserRepo(s Storage) *UserRepo {
	ur := &UserRepo{
		storage: s,
	}
	return ur
}

func (u *UserRepo) Create(ctx context.Context, user userentity.User) (uuid.UUID, error) {
	return u.storage.Create(ctx, user)
}

func (u *UserRepo) Read(ctx context.Context, userId uuid.UUID) (userentity.User, error) {
	return u.storage.Read(ctx, userId)
}

func (u *UserRepo) Update(ctx context.Context, user userentity.User) error {
	return u.storage.Update(ctx, user)
}

func (u *UserRepo) Delete(ctx context.Context, userId uuid.UUID) (userentity.User, error) {
	return u.storage.Delete(ctx, userId)
}

func (u *UserRepo) Search(ctx context.Context, login string) (userentity.User, error) {
	return u.storage.Search(ctx, login)
}
