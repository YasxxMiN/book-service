package routes

import (
	"test-go-book/controllers"
	middlewares "test-go-book/pkg"

	"test-go-book/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.AuthController, repo repositories.AuthRepository) {
	api := app.Group("/api/v1")
	auth := api.Group("/auth")

	auth.Post("/login", userController.Login)
	auth.Get("/auth-test", userController.AuthTest)
	auth.Get("/user-info", middlewares.JwtAuthentication(repo), userController.GetUserInfo)
	api.Patch("/users/me", middlewares.JwtAuthentication(repo), userController.UpdateUserInfo)
	api.Patch("/users/password", middlewares.JwtAuthentication(repo), userController.ChangePassword)
	app.Post("/addbook", middlewares.JwtAuthentication(repo), userController.AddBook)
	app.Delete("/delete", middlewares.JwtAuthentication(repo), userController.DeleteBookUser)
	app.Put("/update-book/:book_id", middlewares.JwtAuthentication(repo), userController.UpdateBookUser)
	app.Get("/get-book", middlewares.JwtAuthentication(repo), userController.GetBookUser)
	auth.Post("/logout", userController.Logout)
}
