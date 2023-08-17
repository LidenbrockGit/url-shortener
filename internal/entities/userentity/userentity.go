package userentity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uuid.UUID
	Name      string
	Login     string
	Password  string
	CreatedAt time.Time
	DeletedAt time.Time
}

func (u *User) IsDeleted() bool {
	return u.DeletedAt != (time.Time{})
}

func (u *User) HashPassword() error {
	bytes, err := HashPassword([]byte(u.Password))
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return CheckPassword([]byte(u.Password), []byte(password))
}

func HashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func CheckPassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}
