package jwt

import (
	"crypto"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// TODO: this is fine?

type Validator struct {
	publicKey crypto.PublicKey
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func NewValidator(publicKeyPEMPath string) (*Validator, error) {
	publicKey, err := ParseEd25519PublicKey(publicKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("read ed25519 public key error: %w", err)
	}

	return &Validator{publicKey}, nil
}

func (m *Validator) Decode(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodEd25519)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return m.publicKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
