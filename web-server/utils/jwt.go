package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a new JWT token for a given email and userId.
func GenerateToken(email string, userId int64) (string, error) {
	// Define the claim for the token, including email, user ID, and the expiration time.
	claims := jwt.MapClaims{
		"email": email,
		"id":    userId,
		"exp":   time.Now().Add(time.Hour * 4).Unix(),
	}
	// Create a new token with the specified signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and return the token as a string using the secret key from the environment variable.
	return token.SignedString([]byte(os.Getenv("JWT_SECRETKEY")))
}

func verifyMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}

	return []byte(os.Getenv("JWT_SECRETKEY")), nil
}

func VerifyToken(tokenStr string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenStr, verifyMethod)

	if err != nil {
		return 0, errors.New("Could not verify token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Token is not valid")
	}

	//type checking parsedToken.Claims.(jwt.MapClaims)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Could not parse claims")
	}

	//type checking claims["email"].(string)
	// email := claims["email"].(string)
	id := int64(claims["id"].(float64))

	return id, nil
}
