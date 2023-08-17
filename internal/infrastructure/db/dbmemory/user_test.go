package dbmemory

import (
	"context"
	"testing"
	"time"

	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/google/uuid"
)

func TestUserCreate(t *testing.T) {
	t.Run("Successful create user", func(t *testing.T) {
		newUser := userentity.User{
			Name:     "Ivan",
			Login:    "Navi",
			Password: "password",
		}
		uRepo := UserRepo{
			memStorage: make(map[uuid.UUID]userentity.User),
		}
		userId, err := uRepo.Create(context.Background(), newUser)
		if err != nil {
			t.Error(err)
			return
		}
		createdUser, ok := uRepo.memStorage[userId]
		if !ok {
			t.Error("cannot find user in storage")
			return
		}
		if createdUser.CreatedAt == (time.Time{}) {
			t.Error("field CreatedAt cannot be empty")
			return
		}
	})

	t.Run("Trying to create duplicate", func(t *testing.T) {
		uRepo, err := createUserRepo()
		if err != nil {
			t.Error(err)
			return
		}

		var newUser userentity.User
		for _, u := range uRepo.memStorage {
			newUser = u
			break
		}

		_, err = uRepo.Create(context.Background(), newUser)
		if err == nil {
			t.Errorf("user with login '%s' cannot be added", newUser.Login)
			return
		}
	})
}

func TestUserRead(t *testing.T) {
	t.Run("Successful read user", func(t *testing.T) {
		uRepo, err := createUserRepo()
		if err != nil {
			t.Error(err)
			return
		}

		userName := "Viktor"
		userId, err := uRepo.Create(context.Background(), userentity.User{Name: userName})
		if err != nil {
			t.Error(err)
			return
		}

		user, err := uRepo.Read(context.Background(), userId)
		if err != nil {
			t.Error(err)
			return
		}

		if user.Name != userName {
			t.Error("wrong user found")
		}
	})
}

func TestUserUpdate(t *testing.T) {
	t.Run("Successful user update", func(t *testing.T) {
		uRepo, err := createUserRepo()
		if err != nil {
			t.Error(err)
			return
		}

		var user, oldUser userentity.User
		for _, v := range uRepo.memStorage {
			user, oldUser = v, v
			break
		}

		if user == (userentity.User{}) {
			t.Error("wrong user found")
			return
		}

		user.Name = "Viktor"
		user.Login = "Viktor"
		user.Password = "123456"
		err = uRepo.Update(context.Background(), user)
		if err != nil {
			t.Error(err)
			return
		}

		if len(uRepo.memStorage) > 1 {
			t.Error("users in memory cannot be more than 1")
			return
		}

		for _, user := range uRepo.memStorage {
			if oldUser.Id != user.Id {
				t.Error("user id shouldn not change")
				return
			}
			if oldUser.Name == user.Name {
				t.Error("user name was not changed")
				return
			}
			if oldUser.Login == user.Login {
				t.Error("user login was not changed")
				return
			}
			if oldUser.Password == user.Password {
				t.Error("user password was not changed")
				return
			}
			break
		}
	})
}

func TestUserDelete(t *testing.T) {
	t.Run("Successful user delete", func(t *testing.T) {
		uRepo, err := createUserRepo()
		if err != nil {
			t.Error(err)
			return
		}

		var user userentity.User
		for _, u := range uRepo.memStorage {
			user = u
			break
		}

		deletedUser, err := uRepo.Delete(context.Background(), user.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if !deletedUser.IsDeleted() {
			t.Error("returned link must be deleted")
			return
		}

		if len(uRepo.memStorage) != 1 {
			t.Error("returned value must be 1")
			return
		}
	})
}

func createUserRepo() (*UserRepo, error) {
	newUser := userentity.User{
		Name:     "Ivan",
		Login:    "Navi",
		Password: "password",
	}

	uRepo := &UserRepo{
		memStorage: make(map[uuid.UUID]userentity.User),
	}

	_, err := uRepo.Create(context.Background(), newUser)
	if err != nil {
		return nil, err
	}

	return uRepo, nil
}
