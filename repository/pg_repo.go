package repository

import "ProductApis/model"

type Repository interface {
	GetDriverDetails(Id string)
	CreateDriver(driver model.Driver)
	GetAllDriverDetails()
	UpdateDriver(Id string)
	DeleteDriver(Id string, driver model.Driver)
	AddNewCar(Id string, car model.Car)
}
