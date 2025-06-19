package utils

import (
	"cadre-management/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

// AccessToken和RefreshToken有效期（小时）
const (
	AccessTokenExpireHours  = 2
	RefreshTokenExpireHours = 24 * 7
)

type Claims struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Role     string `json:"role"` // 添加角色字段
	jwt.StandardClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userid, password, role string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(setting.AppSetting.JwtExptime) * time.Hour)

	claims := Claims{
		userid,
		EncodeMD5(password),
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    setting.AppSetting.JwtIssur,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 生成access token
func GenerateAccessToken(userid, role string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(AccessTokenExpireHours * time.Hour)

	claims := Claims{
		UserID: userid,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    setting.AppSetting.JwtIssur,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 生成refresh token
func GenerateRefreshToken(userid, role string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(RefreshTokenExpireHours * time.Hour)

	claims := RefreshClaims{
		UserID: userid,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    setting.AppSetting.JwtIssur,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 校验refresh token
func ParseRefreshToken(token string) (*RefreshClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*RefreshClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 生成一对token
func GenerateTokenPair(userid, role string) (TokenPair, error) {
	access, err := GenerateAccessToken(userid, role)
	if err != nil {
		return TokenPair{}, err
	}
	refresh, err := GenerateRefreshToken(userid, role)
	if err != nil {
		return TokenPair{}, err
	}
	return TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}
