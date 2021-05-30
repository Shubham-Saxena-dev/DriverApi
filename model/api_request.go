package model

import "github.com/jinzhu/gorm"

type Driver struct {
	gorm.Model

	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   int64  `json:"phone"`
	Car     []Car  `gorm:"embedded"`
}

type Car struct {
	Color        string  `json:"color"`
	LicensePlate string  `json:"licensePlate"`
	Company      Company `gorm:"embedded"`
}

type Company struct {
	Name     string `json:"name"`
	CarModel string `json:"carModel"`
}
