package repositories

import (
	"errors"
	"fmt"
	"os"
	"test-go-book/entities"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	SignUserAccessToken(user *entities.User) (string, error)
	GetUserByID(userID int) (*entities.User, error)
	UpdateUserInfo(userID int, req *entities.User) error
	ChangePassword(userID int, req *entities.User) error
	AddBookToUser(userID int, reqBook *entities.Book) (entities.User, entities.Book, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) SignUserAccessToken(req *entities.User) (string, error) {
	var userID int
	if err := r.db.Model(&entities.User{}).Where("username = ?", req.Username).Pluck("user_id", &userID).Error; err != nil {
		return "", err
	}

	req.User_ID = userID
	claims := entities.UsersClaims{
		Id:       req.User_ID,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (r *authRepo) GetUserByID(userID int) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, userID).Error; err != nil { 
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) UpdateUserInfo(userID int, req *entities.User) error {
	user := &entities.User{User_ID: userID}

	if err := r.db.Model(user).Updates(map[string]interface{}{
		"name":  req.Name,
		"phone": req.Phone,
		"email": req.Email,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepo) ChangePassword(userID int, req *entities.User) error {
	user := &entities.User{User_ID: userID}
	if err := r.db.Model(user).Updates(map[string]interface{}{
		"password_hash": req.PasswordHash,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepo) AddBookToUser(userID int, reqBook *entities.Book) (entities.User, entities.Book, error) {
	var user entities.User
	var book entities.Book

	userBook := entities.UserBook{
		UserID: userID,
		BookID: reqBook.Book_ID,
	}

	var existingUserBook entities.UserBook
	if err := r.db.Where("user_id = ? AND book_id = ?", userID, reqBook.Book_ID).First(&existingUserBook).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := r.db.Create(&userBook).Error; err != nil {
				return entities.User{}, entities.Book{}, err
			}
		}
	} else {
		fmt.Println("คุณมีรายการนี้แล้ว")
		return user, book, errors.New("duplicate entry")
	}

	if err := r.db.First(&book, reqBook.Book_ID).Error; err != nil {
		return user, book, err
	}

	return user, *reqBook, nil
}