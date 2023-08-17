package routergin

import (
	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/gin-gonic/gin"
)

func getCtxUser(ctx *gin.Context) (user userentity.User) {
	i, exists := ctx.Get("user")
	if !exists {
		return userentity.User{}
	}

	user, ok := i.(userentity.User)
	if !ok {
		return userentity.User{}
	}
	return user
}
