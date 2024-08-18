package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type CustomClaims struct {
	jwt.StandardClaims
	Aud            []string         `json:"aud"`
	Sub            uuid.UUID        `json:"sub"`
	Scope          string           `json:"scope"`
	ResourceAccess map[string]Roles `json:"resource_access"`
}
