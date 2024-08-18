package middleware

import (
	"golang-template-api-service/app/config"
	"golang-template-api-service/app/internal/dto/auth"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Define your public key file path
var publicKeyFile = config.LoadViperConfig().Auth.SecretKeyPath // Replace with the path to your public key PEM file
func AuthMiddlewareWithScopes(requiredScopes ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, err := verifyToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Check if the required scopes are present in the token claims
		if !containsAnyScope(claims, requiredScopes) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Unauthorized User",
			})
		}

		// Pass the user claims and token to the next handler
		c.Locals("user_claims", claims)
		c.Locals("user_id", claims.Sub)
		c.Locals("jwt_token", tokenString)

		return c.Next()
	}
}
func containsAnyScope(claims *auth.CustomClaims, requiredScopes []string) bool {
	scopes := strings.Split(claims.Scope, " ")

	for _, requiredScope := range requiredScopes {
		for _, s := range scopes {
			if s == requiredScope {
				return true
			}
		}
	}

	return false
}

func verifyToken(tokenString string) (*auth.CustomClaims, error) {
	publicKeyBytes, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("error reading public key file: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	token, err := jwt.ParseWithClaims(tokenString, &auth.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*auth.CustomClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	return claims, nil
}
