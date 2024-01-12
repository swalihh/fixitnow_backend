package models

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	User_Name string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Location  string `json:"location"`
	Password  string `json:"password"`
	Otp       int    `json:"otp"`
}
