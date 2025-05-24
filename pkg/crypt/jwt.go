// pkg/crypt/jwt.go
package crypt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	DefaultIssuer = "powergame"
	DefaultSecret = "powergame"
)

var (
	ErrTokenExpired = errors.New("token 已过期")
	ErrTokenInvalid = errors.New("token 无效")
)

type JWT struct {
	secretKey []byte
	issuer    string
}

type CustomClaims struct {
	MerchantId string `json:"merchant_id"`
	UserId     string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWT(secret string, issuer string) *JWT {
	return &JWT{
		secretKey: []byte(secret),
		issuer:    issuer,
	}
}

// 生成 Token
func (j *JWT) Generate(merchantId, userId string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		MerchantId: merchantId,
		UserId:     userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

// 解析 Token
func (j *JWT) Parse(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
