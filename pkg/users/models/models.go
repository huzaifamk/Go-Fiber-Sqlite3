package models

// User is the model for the users table in the database
type User struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment"`
	Fname string `json:"fname" gorm:"type:varchar(250)"`
	Lname string `json:"lname" gorm:"type:varchar(250)"`
	Email string `json:"email" gorm:"type:varchar(250)"`
	Phone int64  `json:"phone" gorm:"type:int"`
}

// TableName sets the insert table name for this struct type
func (b *User) TableName() string {
	return "users"
}
