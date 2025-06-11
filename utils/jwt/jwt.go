package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySecret = []byte("xiatianxiatianqiaoqiaoguoqu")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, username string) (string, error) {
	claims := &MyClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "bluebell",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySecret)
}

func ParseToken(token string) (*MyClaims, error) {
	myClaims := &MyClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, myClaims, func(*jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if parsedToken.Valid {
		return myClaims, nil
	}
	return nil, errors.New("invalid token")
}
