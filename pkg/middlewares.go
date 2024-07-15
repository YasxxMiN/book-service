package middlewares

import (
	"os"
	"test-go-book/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JwtAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or invalid token",
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

		book := &entities.Book{
			Book_ID: claims.Book_id,
		}
		c.Locals("user", user)
		c.Locals("book", book)
		return c.Next()
	}
}


