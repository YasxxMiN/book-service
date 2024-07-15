package entities

type Token struct {
	Token_ID int    `gorm:"type:integer;primary_key"`
	Token    string `gorm:"type:text"`
	UserID   int    `gorm:"type:integer"`
}
