package handlers

import (
	"context"

	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/google/uuid"
)

func (h *Handlers) Regist(user userentity.User) (uuid.UUID, error) {
	err := user.HashPassword()
	if err != nil {
		//TODO: ошибку записать в лог; в ответ выдать другую
		return uuid.UUID{}, err
	}
	return h.Userrepo.Create(context.Background(), user)
}

func (h *Handlers) Login(login, password string) (*userentity.User, error) {
	user, err := h.Userrepo.Search(context.Background(), login)
	if err != nil {
		//TODO: ошибку записать в лог; в ответ выдать другую
		return nil, err
	}

	if err = user.CheckPassword(password); err != nil {
		//TODO: ошибку записать в лог
		return nil, WrongPasswordErr
	}

	return &user, nil
}
