package fullurlsearch

import "github.com/LidenbrockGit/url-shortener/internal/entities/linkentity"

type FullUrlFinder interface {
	FullUrlFind(string) (linkentity.Link, error)
}

type FullUrl struct {
	storage FullUrlFinder
}

func NewFullUrl(s FullUrlFinder) FullUrl {
	fu := FullUrl{
		storage: s,
	}
	return fu
}

func (fu *FullUrl) FindByShort(shortUrl string) (linkentity.Link, error) {
	return fu.storage.FullUrlFind(shortUrl)
}
