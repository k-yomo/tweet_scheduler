package jwt_generator

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var Key = os.Getenv("JWT_SIGNING_KEY")

func GenerateJwt(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(Key))
}
