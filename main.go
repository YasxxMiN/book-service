package main

import (
	"fmt"
	"log"

	configs "test-go-book/config"
	"test-go-book/controllers"
	"test-go-book/entities"
	"test-go-book/repositories"
	"test-go-book/routes"
	"test-go-book/usecases"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := configs.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
		return
	}

	db.AutoMigrate(&entities.User{}, &entities.Book{}, &entities.Token{})

	app := fiber.New()
	authRepo := repositories.NewAuthRepository(db)
	userUsecase := usecases.NewAuthUsecase(authRepo)
	authRouter := app.Group("/auth")
	authController := controllers.NewAuthController(authRouter, userUsecase, db, authRepo)

	routes.SetupRoutes(app, authController, authRepo)

	log.Fatal(app.Listen(":3000"))

}
