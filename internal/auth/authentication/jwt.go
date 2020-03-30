package authentication


import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const secretEnv = "JWT_SECRET"

var secret []byte

func Setup(key string) {
	secret = []byte(key)
}

func MustSigned(claims jwt.MapClaims) string {
	uToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := uToken.SignedString(secret)
	if err != nil {
		panic(err)
	}
	return token
}

func provideKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		return secret, nil
	} else {
		return nil, fmt.Errorf("expect token signed with HMAC but got %v", token.Header["alg"])
	}
}

func ValidateToken(uToken string) (*jwt.Token, error) {

	if len(uToken) == 0 {
		return nil, fmt.Errorf("token length is zero")
	}

	token, err := jwt.Parse(uToken, provideKey)

	if err != nil {
		return nil, err
	}

	return token, nil
}