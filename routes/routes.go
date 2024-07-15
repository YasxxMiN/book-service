package routes

import (
	"test-go-book/controllers"
	middlewares "test-go-book/pkg"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.AuthController) {
	api := app.Group("/api/v1")
	auth := api.Group("/auth")

	auth.Post("/login", userController.Login)
	auth.Get("/auth-test", userController.AuthTest)
	auth.Get("/user-info", middlewares.JwtAuthentication(), userController.GetUserInfo)
	api.Patch("/users/me", middlewares.JwtAuthentication(), userController.UpdateUserInfo)
	api.Patch("/users/password", middlewares.JwtAuthentication(), userController.ChangePassword)
	app.Post("/addbook", middlewares.JwtAuthentication(), userController.AddBook)
}
