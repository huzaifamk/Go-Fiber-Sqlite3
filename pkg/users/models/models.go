package models

type User struct {
	ID    int64
	FName string
	LName string
	Email string
	Phone int64
}

func (b *User) TableName() string {
	return "users"
}
