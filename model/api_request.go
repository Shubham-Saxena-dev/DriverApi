package model

import "github.com/jinzhu/gorm"

type Driver struct {
	gorm.Model

	id      int64
	name    string
	address string
	phone  	int64
	car     []Car
}

type Car struct {
	gorm.Model

	model        string
	color        string
	licensePlate string
	company      Company
}

type Company struct {
	gorm.Model

	name     string
	carModel string
}
