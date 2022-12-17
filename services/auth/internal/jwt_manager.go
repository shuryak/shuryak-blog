package internal

import (
	"auth/internal/entity"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTManager struct {
	privatePemKeyPath string
	publicPemKeyPath  string
	tokenDuration     time.Duration
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func NewJWTManager(privatePemKeyPath, publicPemKeyPath string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{privatePemKeyPath, publicPemKeyPath, tokenDuration}
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

	var Ed25519SigningMethod jwt.SigningMethodEd25519

	jwt.RegisterSigningMethod(Ed25519SigningMethod.Alg(), func() jwt.SigningMethod { return &Ed25519SigningMethod })

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	//token := jwt.NewWithClaims(&Ed25519SigningMethod, claims)

	privateKey, err := ParseEd25519PrivateKey(m.privatePemKeyPath)
	if err != nil {
		return "", fmt.Errorf("read ed25519 private key error: %w", err)
	}

	jwtString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token from private key pem error: %w", err)
	}

	return jwtString, nil
}

func (m *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	publicKey, err := ParseEd25519PublicKey(m.publicPemKeyPath)
	if err != nil {
		return nil, fmt.Errorf("read ed25519 private key error: %w", err)
	}

	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodEd25519)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return publicKey, nil
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
