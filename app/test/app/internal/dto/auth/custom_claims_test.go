package auth

import (
	"golang-template-api-service/app/internal/dto/auth"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestCustomClaims(t *testing.T) {
	// Create a CustomClaims instance
	claims := auth.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer:    "example.com",
		},
		Aud:   []string{"audience1", "audience2"},
		Scope: "read write",
		ResourceAccess: map[string]auth.Roles{
			"resource1": {Roles: []string{"role1", "role2"}},
			"resource2": {Roles: []string{"role3", "role4"}},
		},
	}

	// Check the fields of the claims
	if claims.ExpiresAt <= 0 {
		t.Error("ExpiresAt should be a positive Unix timestamp")
	}

	if claims.Issuer != "example.com" {
		t.Error("Issuer should be 'example.com'")
	}

	if len(claims.Aud) != 2 || claims.Aud[0] != "audience1" || claims.Aud[1] != "audience2" {
		t.Error("Aud should contain 'audience1' and 'audience2'")
	}

	if claims.Scope != "read write" {
		t.Error("Scope should be 'read write'")
	}

	if len(claims.ResourceAccess) != 2 {
		t.Error("ResourceAccess should have 2 entries")
	}

	if len(claims.ResourceAccess["resource1"].Roles) != 2 {
		t.Error("Resource1 should have 2 roles")
	}

	if len(claims.ResourceAccess["resource2"].Roles) != 2 {
		t.Error("Resource2 should have 2 roles")
	}
}
