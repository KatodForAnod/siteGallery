package view

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userId int64) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenStr string) (*jwt.Token, error) { // TODO func name verify?? but it only extract token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUserIdFromToken(token *jwt.Token) (int64, error) {
	claims := token.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("wrong claims data")
	}

	return int64(userId), nil
}
