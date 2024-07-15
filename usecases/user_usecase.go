package usecases

import (
	"test-go-book/entities"
	"test-go-book/repositories"
)

type AuthUsecase interface {
	Login(user *entities.User) (string, error)
	GetUserInfo(userID int) (*entities.User, error)
	UpdateUserInfo(userID int, update *entities.User) error
	ChangePassword(userID int, update *entities.User) error
	AddBookToUser(userID int, req *entities.Book) (entities.User, entities.Book, error)
	DeleteBookUser (userID int, req *entities.Book) error
	UpdateBookUser (userID int, req *entities.Book,bookID string) error
}

type authUsecase struct {
	authRepo repositories.AuthRepository
}

func NewAuthUsecase(authRepo repositories.AuthRepository) AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (u *authUsecase) Login(user *entities.User) (string, error) {
	return u.authRepo.SignUserAccessToken(user)
}

func (u *authUsecase) GetUserInfo(userID int) (*entities.User, error) {
	return u.authRepo.GetUserByID(userID)
}

func (u *authUsecase) UpdateUserInfo(userID int, update *entities.User) error {
	return u.authRepo.UpdateUserInfo(userID, update)
}

func (u *authUsecase) ChangePassword(userID int, update *entities.User) error {
	return u.authRepo.ChangePassword(userID, update)
}

func (u *authUsecase) AddBookToUser(userID int, req *entities.Book) (entities.User, entities.Book, error) {
	return u.authRepo.AddBookToUser(userID, req)
}

func (u *authUsecase) DeleteBookUser (userID int, req *entities.Book) error{
	return u.authRepo.DeleteBookUser(userID,req)
}

func (u *authUsecase) UpdateBookUser (userID int, req *entities.Book,bookID string) error {
	return u.authRepo.UpdateBookUser(userID,req,bookID)
}
