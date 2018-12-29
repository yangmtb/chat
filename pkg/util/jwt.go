package util

import (
	"chat/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
)

// Claims ...
type Claims struct {
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken ...
func GenerateToken(username, password string) (token string, err error) {
	claims := Claims{
		username,
		//Sha256String(username),
		//password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			Issuer:    "chat",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

// ParseToken ...
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if nil != tokenClaims {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
