package models

type User struct {
	CommonModelFields
	UserName string `gorm:"unique; default:null" json:"user_name"`
	Email    string `gorm:"unique" json:"email"`
	MobileNo string `gorm:"unique; default:null" json:"mobile_no"`
	Password string `json:"password"`
}
