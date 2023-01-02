package models

type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (b *User) TableName() string {
	return "users"
}