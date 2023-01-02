package models

type User struct {
	ID    int64
	Fname string
	Lname string
	Email string
	Phone int64
}

func (b *User) TableName() string {
	return "users"
}
