package userauthenticator

import (
	"fmt"
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type UserAuthenticator struct {
	jwt *jwt.Token
}

func NewUserAuthenticator() *UserAuthenticator {
	return &UserAuthenticator{
		jwt: jwt.New(jwt.SigningMethodHS256),
	}
}

func (ua *UserAuthenticator) EncodeUser(userID string) (string, error) {
	err := godotenv.Load("/home/hyvinhiljaa/chat-app-vue/server/.env")
	if err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is missing in .env file")
	}

	claims := ua.jwt.Claims.(jwt.MapClaims)
	claims["userID"] = userID

	tokenString, err := ua.jwt.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ua *UserAuthenticator) DecodeUser(tokenString string) (string, error) {
	err := godotenv.Load("/home/hyvinhiljaa/chat-app-server/.env")
	if err != nil {
		return "", fmt.Errorf("failed to load .env file: %w", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET is missing in .env file")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to parse token claims")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse userID from token claims")
	}

	return userID, nil
}
