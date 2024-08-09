package Helper

import (
	"BearLibrary/Config"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

func CreateToken(role uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.Secret))
}

func GetClaims(reqToken string) string {
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 1 {
		return ""
	}
	reqToken = splitToken[1]
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.Secret), nil
	})
	if err != nil {
		return ""
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["role"].(string)
}
