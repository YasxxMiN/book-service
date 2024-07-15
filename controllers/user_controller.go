package controllers

import (
	"test-go-book/entities"
	middlewares "test-go-book/pkg"

	"test-go-book/usecases"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthUsecase usecases.AuthUsecase
}

func NewAuthController(r fiber.Router, authUse usecases.AuthUsecase) *AuthController {
	controller := &AuthController{
		AuthUsecase: authUse,
	}

	r.Post("/login", controller.Login)
	r.Get("/auth-test", middlewares.JwtAuthentication(), controller.AuthTest)
	return controller
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	token, err := controller.AuthUsecase.Login(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Login failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (controller *AuthController) AuthTest(c *fiber.Ctx) error {
	return c.SendString("Auth Test Successful")
}

func (controller *AuthController) GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)
	userInfo, err := controller.AuthUsecase.GetUserInfo(user.User_ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching user info",
		})
	}
	return c.Status(fiber.StatusOK).JSON(userInfo)
}

func (controller *AuthController) UpdateUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)

	var updateRequest entities.User
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := controller.AuthUsecase.UpdateUserInfo(user.User_ID, &updateRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user info",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User info updated successfully",
	})
}

func (controller *AuthController) ChangePassword(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)

	var changeReq entities.User
	if err := c.BodyParser(&changeReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := controller.AuthUsecase.ChangePassword(user.User_ID, &changeReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user info",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User info change password successfully",
	})

}

func (controller *AuthController) AddBook(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)
	var request entities.Book

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user1, book1, err := controller.AuthUsecase.AddBookToUser(user.User_ID, &request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "book added to user successfully",
		"user":    user1.Username,
		"book":    book1.Title,
	})
}

func (controller *AuthController) DeleteBookUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)
	var request entities.Book

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := controller.AuthUsecase.DeleteBookUser(user.User_ID, &request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Delete book successfully",
	})

}

func (controller *AuthController) UpdateBookUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)
	bookID := c.Params("book_id")
	var request entities.Book

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := controller.AuthUsecase.UpdateBookUser(user.User_ID, &request, bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Update book successfully",
	})
}

func (controller *AuthController) GetBookUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*entities.User)
	mybooks, err := controller.AuthUsecase.GetBookUser(user.User_ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "con not get your books",
		})
	}
	return c.Status(fiber.StatusOK).JSON(mybooks)
}
