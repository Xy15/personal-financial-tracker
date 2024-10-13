package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRes struct {
	NewAccessToken string `json:"new_access_token"`
}

type RefreshTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type AccessTokenClaims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewToken(claims jwt.Claims) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return &signedToken, err
}

func GenerateAccessToken(userID string, email string, username string) (*string, error) {
	accessClaims := AccessTokenClaims{
		UserID:   userID,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	signedAccessToken, err := NewToken(accessClaims)

	return signedAccessToken, err
}

func GenerateRefreshToken(userID string) (*string, error) {
	refreshClaims := RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}

	signedRefreshToken, err := NewToken(refreshClaims)

	return signedRefreshToken, err
}

func ParseAccessToken(accessToken string) (*AccessTokenClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	} else if !parsedAccessToken.Valid {
		return nil, errors.New("invalid Access Token")
	}

	return parsedAccessToken.Claims.(*AccessTokenClaims), nil
}

func ParseRefreshToken(refreshToken string) (*RefreshTokenClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	} else if !parsedRefreshToken.Valid {
		return nil, errors.New("invalid Refresh Token")
	}

	return parsedRefreshToken.Claims.(*RefreshTokenClaims), nil
}
