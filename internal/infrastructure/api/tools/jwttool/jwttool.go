package jwttool

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT creates and returns json web token using the secret key.
// The function uses HS256 algorithm
func GenerateJWT(payload map[string]interface{}, secret []byte) (string, error) {
	mc := jwt.MapClaims(payload)
	mc["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), mc)

	return token.SignedString(secret)
}

// ParseJWT returned parsed to string payload
// The function uses HS256 algorithm
func ParseJWT(tokenString string, secret []byte) (map[string]interface{}, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot get payload")
	}

	exp, ok := payload["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return nil, errors.New("time expired")
	}

	return payload, nil
}
