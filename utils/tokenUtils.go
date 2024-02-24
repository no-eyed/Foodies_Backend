package utils

import (
	"crypto/rsa"
	"encoding/pem"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GetSessionIDFromToken(tokenString string) (string, error) {
	claims := &CustomClaims{}

	var PEM_PUBLIC_KEY = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA95S0DeZtfTErA4FGneAU
M20L59mX5qztD8QqakXXbri9/5OC/ZhPv7JT94oz9yHHiHlrcUKmBgfQVUV18iKZ
zRNVQb6KYsknrxJmIFAQvGaazsoyVOewfBeqHU5kVq5aSM9lG8euEEtDNUNFSC74
1bp3m9MEmyCXgQ5q3s0Yg5chGdDrv64RxtxmI0Nx7pZXHEWvUk9oimVrh391Y5eo
5zv0P6Fubnfxrv/ytNwJRofYx4ynyAub+CnQRbq+AP2RrF7IYcr8NK4/ujVB95VL
zjeAAISO3A8L8bVzDDWf8Uw1aAkThUYPjcv5nCiGhfRZ5bAXMy9MP/ueG8QF67yT
wwIDAQAB
-----END PUBLIC KEY-----`

	publicKey, err := parseRSAPublicKeyFromPEM(PEM_PUBLIC_KEY)
	if err != nil {
		return "", fmt.Errorf("error parsing RSA public key: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("error decoding JWT: %v", err)
	}

	if token.Valid {
		return claims.Sid, nil
	}

	return "", fmt.Errorf("invalid token")

}

type CustomClaims struct {
	Sid string `json:"sid"`
	jwt.StandardClaims
}

func parseRSAPublicKeyFromPEM(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}
	pub, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}
	return pub, nil
}
