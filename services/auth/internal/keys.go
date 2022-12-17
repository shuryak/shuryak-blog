package internal

import (
	"crypto"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

// https://blainsmith.com/articles/signing-jwts-with-gos-crypto-ed25519/
// https://engineering.getweave.com/post/go-jwt-1/

func ParseEd25519PrivateKey(privateKeyPEMPath string) (crypto.PrivateKey, error) {
	if privateKeyPEMPath == "" {
		return nil, fmt.Errorf("no ed25519 private key given")
	}

	privateKeyPEM, err := os.ReadFile(privateKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("no ed25519 private key found")
	}

	privateKey, err := jwt.ParseEdPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("ed25519 private key parse error")
	}

	return privateKey, nil
}

func ParseEd25519PublicKey(publicKeyPEMPath string) (crypto.PublicKey, error) {
	if publicKeyPEMPath == "" {
		return nil, fmt.Errorf("no ed25519 public key given")
	}

	publicKeyPEM, err := os.ReadFile(publicKeyPEMPath)
	if err != nil {
		return nil, fmt.Errorf("no ed25519 public key found")
	}

	publicKey, err := jwt.ParseEdPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("ed25519 public key parse error")
	}

	return publicKey, nil
}
