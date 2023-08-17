package linkentity

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	ShortUrl   string
	FullUrl    string
	UsageCount int
	CreatedAt  time.Time
	DeletedAt  time.Time
}

func (l *Link) IsDeleted() bool {
	return l.DeletedAt != (time.Time{})
}
