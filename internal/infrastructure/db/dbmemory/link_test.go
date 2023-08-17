package dbmemory

import (
	"context"
	"testing"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/google/uuid"
)

func TestLinkCreate(t *testing.T) {
	t.Run("Successful create link", func(t *testing.T) {
		newLink := linkentity.Link{
			UserId:   uuid.UUID{},
			ShortUrl: "example",
			FullUrl:  "http://example.com",
		}
		lRepo := LinkRepo{
			memStorage: make(map[uuid.UUID]linkentity.Link),
		}
		link, err := lRepo.Create(context.Background(), newLink)
		if err != nil {
			t.Error(err)
			return
		}
		createdLink, ok := lRepo.memStorage[link.Id]
		if !ok {
			t.Error("cannot find link in storage")
			return
		}
		if createdLink.CreatedAt == (time.Time{}) {
			t.Error("field CreatedAt cannot be empty")
			return
		}
	})

	t.Run("Trying to create duplicate", func(t *testing.T) {
		lRepo, err := createLinkRepo()
		if err != nil {
			t.Error(err)
			return
		}

		var newLink linkentity.Link
		for _, l := range lRepo.memStorage {
			newLink = l
			break
		}

		_, err = lRepo.Create(context.Background(), newLink)
		if err == nil {
			t.Errorf("link with ShortUrl '%s' cannot be added", newLink.ShortUrl)
			return
		}
	})
}

func TestLinkRead(t *testing.T) {
	t.Run("Successful read link", func(t *testing.T) {
		lRepo, err := createLinkRepo()
		if err != nil {
			t.Error(err)
			return
		}

		shortUrl := "example2"
		link, err := lRepo.Create(context.Background(), linkentity.Link{ShortUrl: shortUrl})
		if err != nil {
			t.Error(err)
			return
		}

		link, err = lRepo.Read(context.Background(), link.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if link.ShortUrl != shortUrl {
			t.Error("wrong link found")
		}
	})
}

func TestLinkReadAll(t *testing.T) {
	t.Run("Successful read all links", func(t *testing.T) {
		lRepo, err := createLinkRepo()
		if err != nil {
			t.Error(err)
			return
		}

		_, err = lRepo.Create(context.Background(), linkentity.Link{
			UserId:   uuid.New(),
			ShortUrl: "example2",
			FullUrl:  "http://example2.com",
		})
		if err != nil {
			t.Error(err)
			return
		}

		chIn, err := lRepo.ReadAll(context.Background())
		if err != nil {
			t.Error(err)
			return
		}

		var links []linkentity.Link
		for l := range chIn {
			links = append(links, l)
		}

		if len(links) != 2 {
			t.Error("links len must be 2")
			return
		}
	})
}

func TestLinkUpdate(t *testing.T) {
	t.Run("Successful link update", func(t *testing.T) {
		lRepo, err := createLinkRepo()
		if err != nil {
			t.Error(err)
			return
		}

		var link, oldLink linkentity.Link
		for _, l := range lRepo.memStorage {
			link, oldLink = l, l
			break
		}

		if link == (linkentity.Link{}) {
			t.Error("wrong link found")
			return
		}

		link.UserId = uuid.New()
		link.ShortUrl = "example2"
		link.FullUrl = "http://example2.com"
		err = lRepo.Update(context.Background(), link)
		if err != nil {
			t.Error(err)
			return
		}

		if len(lRepo.memStorage) > 1 {
			t.Error("links in memory cannot be more than 1")
			return
		}

		for _, link := range lRepo.memStorage {
			if oldLink.Id != link.Id {
				t.Error("link id shouldn not change")
				return
			}
			if oldLink.ShortUrl == link.ShortUrl {
				t.Error("link name was not changed")
				return
			}
			if oldLink.FullUrl == link.FullUrl {
				t.Error("link login was not changed")
				return
			}
			break
		}
	})
}

func TestLinkDelete(t *testing.T) {
	t.Run("Successful link delete", func(t *testing.T) {
		lRepo, err := createLinkRepo()
		if err != nil {
			t.Error(err)
			return
		}

		var link linkentity.Link
		for _, l := range lRepo.memStorage {
			link = l
			break
		}

		deletedLink, err := lRepo.Delete(context.Background(), link.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if !deletedLink.IsDeleted() {
			t.Error("returned link must be deleted")
			return
		}

		if len(lRepo.memStorage) != 1 {
			t.Error("returned value must be 1")
			return
		}
	})
}

func createLinkRepo() (*LinkRepo, error) {
	newLink := linkentity.Link{
		UserId:   uuid.New(),
		ShortUrl: "example",
		FullUrl:  "http://example.com",
	}

	lRepo := &LinkRepo{
		memStorage: make(map[uuid.UUID]linkentity.Link),
	}

	_, err := lRepo.Create(context.Background(), newLink)
	if err != nil {
		return nil, err
	}

	return lRepo, nil
}
