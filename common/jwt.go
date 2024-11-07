package common

import "github.com/golang-jwt/jwt/v5"

// JwtCustomClaims custom claims
type JwtCustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Roles    string `json:"roles"`
	jwt.RegisteredClaims
}
