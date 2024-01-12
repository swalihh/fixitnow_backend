package models

type Saved struct {
	Id          int `json:"id" gorm:"primaryKey"`
	User_Id     int `json:"userid"`
	Servicer_Id int `json:"service_id"`
}
