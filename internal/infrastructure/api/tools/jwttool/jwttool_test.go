package jwttool

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	t.Run("Successful generate jwt", func(t *testing.T) {
		payload := make(map[string]interface{})
		payload["userId"] = 1

		tokenStr, err := GenerateJWT(payload, []byte("secret"))
		if !assert.NoError(t, err) {
			return
		}

		token, err := jwt.Parse(
			tokenStr,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			},
		)
		if !assert.NoError(t, err) {
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !assert.True(t, ok) {
			return
		}

		userId, ok := payload["userId"].(float64)
		if !assert.True(t, ok) {
			return
		}

		if !assert.Equal(t, userId, float64(1), "userId must be 1") {
			return
		}
	})
}

func TestParseJWT(t *testing.T) {
	t.Run("Successful parse jwt", func(t *testing.T) {
		// Generate jwt
		payload := make(map[string]interface{})
		payload["userId"] = 1
		payload["exp"] = time.Now().Add(time.Hour * 24).Unix()
		mc := jwt.MapClaims(payload)
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), mc)
		tokenStr, err := token.SignedString([]byte("secret"))
		if !assert.NoError(t, err) {
			return
		}

		// testing ParseJWT
		payload, err = ParseJWT(tokenStr, []byte("secret"))
		if !assert.NoError(t, err) {
			return
		}

		userId, ok := payload["userId"].(float64)
		if !assert.True(t, ok) {
			return
		}

		if !assert.Equal(t, userId, float64(1), "userId must be 1") {
			return
		}
	})

	t.Run("time expired", func(t *testing.T) {
		// Generate jwt
		payload := make(map[string]interface{})
		payload["userId"] = 1
		payload["exp"] = time.Now().Add(-time.Second).Unix()
		mc := jwt.MapClaims(payload)
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), mc)
		tokenStr, err := token.SignedString([]byte("secret"))
		if !assert.NoError(t, err) {
			return
		}

		// testing ParseJWT
		_, err = ParseJWT(tokenStr, []byte("secret"))
		if !assert.Error(t, err) {
			return
		}
	})
}
