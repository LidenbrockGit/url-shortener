package auth

import (
	"net/http"
	"strings"

	"github.com/LidenbrockGit/url-shortener/internal/entities/userentity"
	"github.com/LidenbrockGit/url-shortener/internal/infrastructure/api/tools/jwttool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserSearch func(userId string) (userentity.User, error)

func GinAuthMW(ctx *gin.Context, us UserSearch) {
	ok := func() bool {
		jwtKey := []byte(viper.GetString("jwt_key"))
		payload, err := jwttool.ParseJWT(getAuthToken(ctx), jwtKey)
		if err != nil {
			return false
		}

		userId, ok := payload["userId"].(string)
		if !ok {
			return false
		}

		user, err := us(userId)
		if err != nil {
			return false
		}

		ctx.Set("user", user)
		return true
	}()
	if !ok {
		// TODO: записать в логи err
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "wrong JWT token"})
		ctx.Abort()
		return
	}
}

func getAuthToken(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	authHeaderParts := strings.Split(authHeader, "Bearer ")
	if len(authHeaderParts) < 2 {
		return ""
	}
	return authHeaderParts[1]
}
