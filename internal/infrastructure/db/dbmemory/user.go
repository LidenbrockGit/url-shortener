package dbmemory

import (
	"context"
	"fmt"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/LidenbrockGit/url-shortener/internal/usecases/repos/userrepo"
	"github.com/google/uuid"
)

var _ userrepo.Storage = &UserRepo{}

type UserRepo struct {
	memStorage map[uuid.UUID]userentity.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		memStorage: make(map[uuid.UUID]userentity.User),
	}
}

func (db *UserRepo) Create(ctx context.Context, user userentity.User) (uuid.UUID, error) {
	select {
	case <-ctx.Done():
		return uuid.UUID{}, ctx.Err()
	default:
	}

	for _, u := range db.memStorage {
		if !u.IsDeleted() && u.Login == user.Login {
			return uuid.UUID{}, fmt.Errorf("%w; login '%s'", UserLoginDuplicatingErr, user.Login)
		}
	}

	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	db.memStorage[user.Id] = user

	return user.Id, nil
}

func (db *UserRepo) Read(ctx context.Context, userId uuid.UUID) (userentity.User, error) {
	select {
	case <-ctx.Done():
		return userentity.User{}, ctx.Err()
	default:
	}

	user, ok := db.memStorage[userId]
	if !ok || user.IsDeleted() {
		return userentity.User{}, fmt.Errorf("%w; userId '%s'", UserFindErr, userId)
	}

	return user, nil
}

func (db *UserRepo) Update(ctx context.Context, user userentity.User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err := db.Read(ctx, user.Id)
	if err != nil {
		return err
	}

	db.memStorage[user.Id] = user

	return nil
}

func (db *UserRepo) Delete(ctx context.Context, userId uuid.UUID) (userentity.User, error) {
	select {
	case <-ctx.Done():
		return userentity.User{}, ctx.Err()
	default:
	}

	deleteUser, err := db.Read(ctx, userId)
	if err != nil {
		return userentity.User{}, err
	}

	deleteUser.DeletedAt = time.Now()
	err = db.Update(ctx, deleteUser)
	if err != nil {
		return userentity.User{}, err
	}

	return deleteUser, nil
}

func (db *UserRepo) Search(ctx context.Context, login string) (userentity.User, error) {
	select {
	case <-ctx.Done():
		return userentity.User{}, ctx.Err()
	default:
	}

	for _, user := range db.memStorage {
		if user.Login == login {
			return user, nil
		}
	}
	return userentity.User{}, UserFindErr
}
