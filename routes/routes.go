package routes

import (
	"test-go-book/controllers"
	middlewares "test-go-book/pkg"

	"test-go-book/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.AuthController, repo repositories.AuthRepository) {

	app.Post("/login", userController.Login)
	app.Get("/auth-test", userController.AuthTest)
	app.Get("/user-info", middlewares.JwtAuthentication(repo), userController.GetUserInfo)
	app.Patch("/users/me", middlewares.JwtAuthentication(repo), userController.UpdateUserInfo)
	app.Patch("/users/password", middlewares.JwtAuthentication(repo), userController.ChangePassword)
	app.Post("/addbook", middlewares.JwtAuthentication(repo), userController.AddBook)
	app.Delete("/delete", middlewares.JwtAuthentication(repo), userController.DeleteBookUser)
	app.Put("/update-book/:book_id", middlewares.JwtAuthentication(repo), userController.UpdateBookUser)
	app.Get("/get-book", middlewares.JwtAuthentication(repo), userController.GetBookUser)
	app.Post("/logout", userController.Logout)

}
