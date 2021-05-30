package service

import (
	"ProductApis/model"
	"github.com/jinzhu/gorm"
)

type RouteService interface {
	GetDriverDetails(id int)
	CreateDriver(driver model.Driver)
	GetAllDriverDetails()
	UpdateDriver(id int)
	DeleteDriver(id int)
	AddNewCar(car model.Car)
}

type routeService struct {
	repo *gorm.DB
}

func NewRouteService(repo *gorm.DB) RouteService {
	return &routeService{repo: repo}
}

func (r routeService) GetDriverDetails(id int) {
	panic("implement me")
}

func (r routeService) CreateDriver(driver model.Driver) {
	panic("implement me")
}

func (r routeService) GetAllDriverDetails() {
	panic("implement me")
}

func (r routeService) UpdateDriver(id int) {
	panic("implement me")
}

func (r routeService) DeleteDriver(id int) {
	panic("implement me")
}

func (r routeService) AddNewCar(car model.Car) {
	panic("implement me")
}
