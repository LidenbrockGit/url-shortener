package handlers

import (
	"context"

	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/google/uuid"
)

func (h *Handlers) UserRead(userIdStr string) (userentity.User, error) {
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return userentity.User{}, err
	}

	return h.Userrepo.Read(context.Background(), userId)
}

func (h *Handlers) UserUpdate(user userentity.User) error {
	return h.Userrepo.Update(context.Background(), user)
}

func (h *Handlers) UserDelete(userIdStr string) (userentity.User, error) {
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return userentity.User{}, IncorrectUserIdErr
	}
	return h.Userrepo.Delete(context.Background(), userId)
}
