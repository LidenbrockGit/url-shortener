package linkentity

import "github.com/google/uuid"

type Link struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	ShortUrl   string
	FullUul    string
	UsageCount int
	CreatedAt  string
}
