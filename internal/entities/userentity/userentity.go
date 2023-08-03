package userentity

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Name      string
	Login     string
	Password  string
	CreatedAt string
	DeletedAt string
}
