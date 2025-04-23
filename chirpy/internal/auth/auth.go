package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	tn := time.Now().UTC()
	exp := tn.Add(expiresIn)

	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(tn),
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   UserID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

//	func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
//		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
//			return []byte(tokenSecret), nil
//		}, nil)
//		if err != nil {
//			return uuid.Nil, err
//		}
//
//		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
//			return uuid.Parse(claims.Subject)
//		}
//
//		return uuid.Nil, fmt.Errorf("invalid token claims")
//	}
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	// Parse the token without specifying claims first
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	// Check if token is valid
	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	// Extract the claims manually as a map
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}

	// Get the subject from the map
	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid subject claim")
	}

	// Parse the UUID
	id, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
