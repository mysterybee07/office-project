package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secret")

type JWTSignedDetails struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateJWTToken(email string, role string) (string, string, error) {

	expirationTime := time.Now().Add(time.Hour * 72).Unix()
	claims := &JWTSignedDetails{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	refreshClaims := &JWTSignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime + 30*24*60*60,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign token: %w", err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign token: %w", err)
	}

	return token, refreshToken, nil
}

func SetToken(ctx *gin.Context, accessToken string, refreshToken string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})
}

func ValidateToken(signedToken string, isRefreshToken bool) (*JWTSignedDetails, error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTSignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*JWTSignedDetails)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if isRefreshToken {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("refresh token expired")
		}
	} else {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("access token expired")
		}
	}
	return claims, nil
}
