package model

import "github.com/jinzhu/gorm"

type Driver struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   int64  `json:"phone"`
	Car     []Car  `gorm:"ForeignKey:DriverID"`
}

type Car struct {
	gorm.Model
	Color        string  `json:"color"`
	LicensePlate string  `json:"licensePlate"`
	Company      Company `gorm:"ForeignKey:CarID"`
	DriverID     uint    `gorm:"column:driver_id"`
}

type Company struct {
	gorm.Model
	Name     string `json:"name"`
	CarModel string `json:"carModel"`
	CarID	uint `gorm:"column:car_id"`
}

func (faculty *Driver) TableName() string {
	return "driver"
}

func (faculty *Car) TableName() string {
	return "car"
}

func (faculty *Company) TableName() string {
	return "company"
}
