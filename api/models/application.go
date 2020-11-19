package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Application struct
type Application struct {
	UserID     int `gorm:"primary_key;not null" json:"user_id"`
	JobOfferID int `gorm:"primary_key;not null" json:"job_offer_id"`
}

// Prepare prepares application to save
func (a *Application) Prepare() {
	a.JobOfferID = 0
	a.UserID = 0
}

// Validate validates data before save
func (a *Application) Validate() error {
	if a.UserID < 0 {
		return errors.New("Required UserID")
	}
	if a.JobOfferID < 0 {
		return errors.New("Required JobOfferID")
	}
	return nil
}

// SaveApplication saves new application to DB
func (a *Application) SaveApplication(db *gorm.DB) (*Application, error) {
	var err error
	err = db.Debug().Model(&Application{}).Create(&a).Error
	if err != nil {
		return &Application{}, err
	}
	return a, nil
}
