package entities

type Book struct {
	Book_ID     int    `gorm:"type:integer;primary_key"`
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Author      string `gorm:"type:varchar(100)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
	UserID      int    `gorm:"type:integer" json:"user_id"`
	Users       []User `gorm:"many2many:user_books;"`
}

type BookandUser struct {
	UserID      int
	BookID      int
	Title       string
	Author      string
	Description string
}

type Mybook interface{
	BookandUser
}
