package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// CustomClaims 自定义的Claim
type CustomClaims struct {
	jwt.RegisteredClaims
}

//JwtUser 需要使用jwtToken的用户都需要实现该方法
type JwtUser interface {
	GetUid() string
}

//CreateToken 生成Jwt的Token
func CreateToken(iss string, user JwtUser, secret string, ttl int64) (tokenString string, err error) {
	claims := CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Second)),
			ID:        user.GetUid(),
			Issuer:    iss, //用来区分不同端的jwt使用
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(secret))
	return
}

//ParseToken 把Token字符串转换为数据
func ParseToken(tokenString string, secret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, nil
}
