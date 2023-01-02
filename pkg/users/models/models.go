package models

type User struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"fname" gorm:"type:varchar(250)"`
	LastName  string `json:"lname" gorm:"type:varchar(250)"`
	Email     string `json:"email" gorm:"type:varchar(250)"`
	Phone     int64  `json:"phone" gorm:"type:int"`
}

func (b *User) TableName() string {
	return "users"
}
