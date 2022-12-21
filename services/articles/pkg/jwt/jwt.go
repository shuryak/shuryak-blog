package jwt

import (
	"auth/internal/entity"
	"crypto"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TODO: this is fine?

type Validator struct {
	publicKey crypto.PublicKey
}

type Issuer struct {
	privateKey    crypto.PrivateKey
	tokenDuration time.Duration
}

type Manager struct {
	Validator
	Issuer
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

func NewIssuer(privateKeyPEMPath string, tokenDuration time.Duration) (*Issuer, error) {
	privateKey, err := ParseEd25519PrivateKey(privateKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("read ed25519 private key error: %w", err)
	}

	return &Issuer{
		privateKey,
		tokenDuration,
	}, nil
}

func NewManager(privateKeyPEMPath, publicKeyPEMPath string, tokenDuration time.Duration) (*Manager, error) {
	issuer, err := NewIssuer(privateKeyPEMPath, tokenDuration)
	if err != nil {
		return nil, fmt.Errorf("private key error: %w", err)
	}

	validator, err := NewValidator(publicKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("public key error: %w", err)
	}

	return &Manager{*validator, *issuer}, nil
}

func (m *Issuer) Generate(user *entity.User) (string, error) {
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
