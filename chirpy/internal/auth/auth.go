package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokenType string

const (
	TokenTypeAccess TokenType = "chirpy-access"
)

func HashPassword(password string) (string, error) {
	hashed_passwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed_passwd), nil
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func MakeJWT(UserID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	// tn := time.Now().UTC()
	// exp := tn.Add(expiresIn)

	// claims := jwt.RegisteredClaims{
	// 	Issuer:    string(TokenTypeAccess),
	// 	IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
	// 	Subject:   UserID.String(),
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    string(TokenTypeAccess),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   UserID.String(),
	})

	return token.SignedString([]byte(tokenSecret))
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claimsStruct, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}

	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
	}

	return id, nil
}

// func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
// 	// Parse the token without specifying claims first
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Validate the alg is what you expect
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(tokenSecret), nil
// 	})
//
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
//
// 	// Check if token is valid
// 	if !token.Valid {
// 		return uuid.Nil, fmt.Errorf("invalid token")
// 	}
//
// 	// Extract the claims manually as a map
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return uuid.Nil, fmt.Errorf("invalid token claims")
// 	}
//
// 	// Get the subject from the map
// 	sub, ok := claims["sub"].(string)
// 	if !ok {
// 		return uuid.Nil, fmt.Errorf("invalid subject claim")
// 	}
//
// 	// Parse the UUID
// 	id, err := uuid.Parse(sub)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
//
// 	return id, nil
// }

// func MakeJWT(
// 	userID uuid.UUID,
// 	tokenSecret string,
// 	expiresIn time.Duration,
// ) (string, error) {
// 	signingKey := []byte(tokenSecret)
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
// 		Issuer:    string(TokenTypeAccess),
// 		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
// 		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
// 		Subject:   userID.String(),
// 	})
// 	return token.SignedString(signingKey)
// }
//
// // ValidateJWT -
// func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
// 	claimsStruct := jwt.RegisteredClaims{}
// 	token, err := jwt.ParseWithClaims(
// 		tokenString,
// 		&claimsStruct,
// 		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
// 	)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
//
// 	userIDString, err := token.Claims.GetSubject()
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
//
// 	issuer, err := token.Claims.GetIssuer()
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
// 	if issuer != string(TokenTypeAccess) {
// 		return uuid.Nil, errors.New("invalid issuer")
// 	}
//
// 	id, err := uuid.Parse(userIDString)
// 	if err != nil {
// 		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
// 	}
// 	return id, nil
// }
