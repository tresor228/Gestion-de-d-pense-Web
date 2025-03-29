package middleware

import (
	"gestion-de-depense/backend/pkg/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware protège les routes avec JWT
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenString := strings.Split(authHeader, "Bearer ")
	if len(tokenString) != 2 {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token format"})
	}

	token, err := auth.ValidateToken(tokenString[1])
	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token claims"})
	}

	c.Locals("user_id", claims["user_id"])
	return c.Next()
}
