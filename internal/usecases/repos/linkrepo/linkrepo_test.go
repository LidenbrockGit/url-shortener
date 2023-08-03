package linkrepo

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"
	"github.com/LidenbrockGit/url-shortener/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {
	t.Run("Storage.ReadAll sends 5 links to channel", func(t *testing.T) {
		chOut := make(chan linkentity.Link)

		str := mocks.NewStorage(t)
		str.EXPECT().ReadAll(context.Background()).Return(chOut, nil)

		go func() {
			for i := 0; i < 5; i++ {
				chOut <- linkentity.Link{}
			}
			close(chOut)
		}()

		linkRepo := LinkRepo{
			storage: str,
		}

		chLinks, err := linkRepo.ReadAll(context.Background())
		if !assert.NotNil(t, chLinks, "chLinks cannot be nil") {
			return
		}
		if !assert.NoError(t, err, "err must be nil") {
			return
		}

		var links []linkentity.Link
		for link := range chLinks {
			links = append(links, link)
		}

		if !assert.Equal(t, 5, len(links), "the number of links must be equal to 5") {
			return
		}
	})

	t.Run("Storage.ReadAll returned error", func(t *testing.T) {
		str := mocks.NewStorage(t)
		str.EXPECT().ReadAll(context.Background()).Return(nil, errors.New("some error"))

		linkRepo := LinkRepo{
			storage: str,
		}

		chLinks, err := linkRepo.ReadAll(context.Background())
		if !assert.Nil(t, chLinks, "chLinks must be nil") {
			return
		}
		if !assert.Error(t, err, "err cannot be nil") {
			return
		}
	})

	t.Run("Context is done", func(t *testing.T) {
		chOut := make(chan linkentity.Link)

		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*100))
		defer cancel()

		str := mocks.NewStorage(t)
		str.EXPECT().ReadAll(ctx).Return(chOut, nil)

		go func() {
			defer close(chOut)
			chOut <- linkentity.Link{}
			<-ctx.Done()
			chOut <- linkentity.Link{}
			time.Sleep(time.Millisecond * 100)
		}()

		linkRepo := LinkRepo{
			storage: str,
		}
		chLinks, err := linkRepo.ReadAll(ctx)
		if !assert.NotNil(t, chLinks, "chLinks cannot be nil") {
			return
		}
		if !assert.NoError(t, err, "err must be nil") {
			return
		}

		linksNumber := 0
		for range chLinks {
			linksNumber++
		}

		if !assert.Equal(t, 1, linksNumber, "number of links must be 1") {
			return
		}
	})

	t.Run("Timer expired", func(t *testing.T) {
		chOut := make(chan linkentity.Link)

		str := mocks.NewStorage(t)
		str.EXPECT().ReadAll(context.Background()).Return(chOut, nil)

		go func() {
			defer close(chOut)
			chOut <- linkentity.Link{}
			time.Sleep(time.Millisecond * 2100)
			chOut <- linkentity.Link{}
		}()

		linkRepo := LinkRepo{
			storage: str,
		}
		chLinks, err := linkRepo.ReadAll(context.Background())
		if !assert.NotNil(t, chLinks, "chLinks cannot be nil") {
			return
		}
		if !assert.NoError(t, err, "err must be nil") {
			return
		}

		linksNumber := 0
		for range chLinks {
			linksNumber++
		}

		if !assert.Equal(t, 1, linksNumber, "number of links must be 1") {
			return
		}
	})
}
