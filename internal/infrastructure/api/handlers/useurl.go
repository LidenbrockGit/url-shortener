package handlers

import "context"

func (h *Handlers) UseUrl(shortUrl string) (fullUrl string, err error) {
	link, err := h.Fullurlsearch.FindByShort(context.Background(), shortUrl)
	if err != nil {
		return "", err
	}

	return link.FullUrl, nil
}
