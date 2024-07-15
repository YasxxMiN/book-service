package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	User_ID      int    `gorm:"type:integer;primary_key"`
	Username     string `gorm:"type:varchar(10);unique_index;not null" json:"username"`
	PasswordHash string `gorm:"type:varchar(10);not null" json:"password_hash"`
	Name         string `gorm:"type:varchar(10)" json:"name"`
	Email        string `gorm:"type:varchar(10);unique_index" json:"email"`
	Phone        string `gorm:"type:varchar(10)" json:"phone"`
	Books        []Book `gorm:"many2many:user_books;"`
}

type MyUser struct {
	UserID   int
	UserName string
	Password string
	Name     string
	Email    string
	Phone    string
}

type UserBook struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}


type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Book_id  int    `json:"book_id"`
	jwt.RegisteredClaims
}

type AuthRepository interface {
	SignUserAccessToken(req *User) (string, error)
}

type UsersRepository interface {
	FindOneUser(username string) (*User, error)
}

type AuthUsecase interface {
	Login(req *User) (*UsersLoginRes, error)
}


