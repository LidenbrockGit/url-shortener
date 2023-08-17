package handlers

import (
	"context"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/google/uuid"
)

func (h *Handlers) CreateLink(link linkentity.Link) (linkentity.Link, error) {
	link.CreatedAt = time.Now()
	return h.Linkrepo.Create(context.Background(), link)
}

func (h *Handlers) GetLinks() (<-chan linkentity.Link, error) {
	return h.Linkrepo.ReadAll(context.Background())
}

func (h *Handlers) GetLink(linkIdStr string) (linkentity.Link, error) {
	linkUUID, err := uuid.Parse(linkIdStr)
	if err != nil {
		return linkentity.Link{}, IncorrectLinkIdErr
	}
	return h.Linkrepo.Read(context.Background(), linkUUID)
}

func (h *Handlers) UpdateLink(link linkentity.Link) error {
	return h.Linkrepo.Update(context.Background(), link)
}

func (h *Handlers) DeleteLink(linkIdStr string) (linkentity.Link, error) {
	linkUUID, err := uuid.Parse(linkIdStr)
	if err != nil {
		return linkentity.Link{}, IncorrectLinkIdErr
	}
	return h.Linkrepo.Delete(context.Background(), linkUUID)
}
