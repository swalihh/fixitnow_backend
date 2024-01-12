package models

type Servicer struct {
	ID                   int    `json:"id" gorm:"primaryKey;not null"`
	UserName             string `json:"username" gorm:"not null"`
	Email                string `json:"email" gorm:"not null"`//
	PhoneNumber          string `json:"phone" gorm:"not null"`
	Password             string `json:"password" gorm:"not null"`//
	OTP                  string `json:"otp" gorm:"not null"`//
	FullName             string `json:"fullname"`
	Description          string `json:"description"`
	ServiceCatagory      string `json:"servicecatagory"`
	VerificationDocument string `json:"verificationdocument"`//
	Amount               int    `json:"amount"`
	Location             string `json:"location"`
	ServicerImage        string `json:"servicerimage"`
	ServicerDocument     string `json:"servicerdocument"`//
	Status               string `json:"status"`
}


