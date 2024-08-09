package Helper

import (
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

func CreateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["authorization"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func GetClaims(reqToken string) bool {
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 1 {
		return false
	}
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return false
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["authorization"].(bool)
}
