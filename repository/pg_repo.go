package repository

import (
	"ProductApis/model"
	"github.com/jinzhu/gorm"
)

const (
	CAR     = "Car"
	Company = "Car.Company"
)

type Repository interface {
	GetDriverDetails(id int) (model.Driver, error)
	CreateDriver(driver model.Driver) error
	GetAllDriverDetails() ([]model.Driver,error)
	UpdateDriver(Id string)
	DeleteDriver(Id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetDriverDetails(id int) (model.Driver, error) {
	var driver model.Driver
	err := r.db.Preload(CAR).Preload(Company).Find(&driver, id).Error
	return driver, err
}

func (r repository) CreateDriver(driver model.Driver) error {
	return r.db.Create(&driver).Error
}

func (r repository) GetAllDriverDetails() ([]model.Driver, error){
	var driver []model.Driver
	err := r.db.Preload(CAR).Preload(Company).Find(&driver).Error
	return driver, err
}

func (r repository) UpdateDriver(Id string) {
	panic("implement me")
}

func (r repository) DeleteDriver(id int) error {
	return r.db.Unscoped().Delete(&model.Driver{}, id).Error
}
