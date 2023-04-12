package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)

	if err != nil {
		return "", fmt.Errorf("Could not decode key %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("Create: parse key %w", err)
	}

	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)


	if err != nil {
		return "", fmt.Errorf("Create: sign token %w", err)
	}
	return token, nil
}

func ValidateToken(token, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)

	if err != nil {
		return "", fmt.Errorf("Could not decode %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return "", fmt.Errorf("Validate: parse key %w", err)
	}

	parsedToken, err := jwt.Parse(token, func (t *jwt.Token) (interface{}, error)  {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Validate: %w", err)
	}


	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("Validate: invalid token")
	}
	return claims["sub"], nil
}