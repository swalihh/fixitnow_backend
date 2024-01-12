package models

import "time"

type Booking struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	BuildingName  string    `json:"buildingname"`
	City          string    `json:"city"`
	Road          string    `json:"road"`
	Phone         string    `json:"phone"`
	Date          string    `json:"date"`
	Time          string    `json:"time"`
	Description   string    `json:"description"`
	Servicer_id   int       `json:"servicerid"`
	ServiceAmount int       `json:"serviceamount"`
	User_Id       int       `json:"userid"`
	Status        string    `json:"status"`
	StartingTime  time.Time `json:"startingtime"`
	EndingTime    time.Time `json:"endingtime"`
	PaymentStatus string    `json:"payment_status" gorm:"default:Pending"`
	Revenue       int64     `json:"revenue" gorm:"default:0"`
}
