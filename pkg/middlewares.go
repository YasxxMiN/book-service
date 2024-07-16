package middlewares

import (
	"os"
	"test-go-book/entities"

	"test-go-book/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JwtAuthentication(authRepo repositories.AuthRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or invalid token",
			})
		}

		blacklisted, err := authRepo.IsTokenBlacklisted(tokenString)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check token validity",
			})
		}
		if blacklisted {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token has been invalidated",
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &entities.UsersClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		claims, ok := token.Claims.(*entities.UsersClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token claims",
			})
		}

		user := &entities.User{
			User_ID:  claims.Id,
			Username: claims.Username,
		}
		c.Locals("user", user)
		return c.Next()
	}
}
