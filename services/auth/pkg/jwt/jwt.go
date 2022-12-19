package internal

import (
	"auth/internal/entity"
	"crypto"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTManager struct {
	privateKey    crypto.PrivateKey
	publicKey     crypto.PublicKey
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func NewJWTManager(privateKeyPEMPath, publicKeyPEMPath string, tokenDuration time.Duration) (*JWTManager, error) {
	privateKey, err := ParseEd25519PrivateKey(privateKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("read ed25519 private key error: %w", err)
	}

	publicKey, err := ParseEd25519PublicKey(publicKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("read ed25519 public key error: %w", err)
	}

	return &JWTManager{privateKey, publicKey, tokenDuration}, nil
}

func (m *JWTManager) Generate(user *entity.User) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
		},
		UserId:   user.Id,
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	jwtString, err := token.SignedString(m.privateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token from private key error: %w", err)
	}

	return jwtString, nil
}

func (m *JWTManager) Verify(accessToken string) (*UserClaims, error) {
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
