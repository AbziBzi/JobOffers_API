package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

// Office structure
type Office struct {
	ID         int    `gorm:"primary_key;auto_increment" json:"id"`
	Country    string `gorm:"size:255;not null" json:"country"`
	City       string `gorm:"size:255;not null" json:"city"`
	ZipCode    string `gorm:"size:255" json:"zip_code"`
	Street     string `gorm:"size:255;not null" json:"street"`
	BuildingNr string `gorm:"size:255" json:"building_nr"`
	CompanyID  int    `gorm:"not null; unique" json:"company_id"`
}

// Prepare func sets office attributes to null
func (o *Office) Prepare() {
	o.ID = 0
	o.Country = strings.TrimSpace(o.Country)
	o.City = strings.TrimSpace(o.City)
	o.ZipCode = strings.TrimSpace(o.ZipCode)
	o.Street = strings.TrimSpace(o.Street)
	o.BuildingNr = strings.TrimSpace(o.BuildingNr)
}

// Validate func checks if given data is valid
func (o *Office) Validate() error {
	if o.Country == "" {
		return errors.New("Required Country")
	}
	if o.City == "" {
		return errors.New("Required City")
	}
	if o.Street == "" {
		return errors.New("Required Street")
	}
	if o.CompanyID < 1 {
		return errors.New("Required CompanyID")
	}
	return nil
}

// SaveOffice func saves office to DB
func (o *Office) SaveOffice(db *gorm.DB) (*Office, error) {
	var err error
	err = db.Debug().Model(&Office{}).Create(&o).Error
	if err != nil {
		return &Office{}, err
	}
	return o, nil
}

// FindAllOffices returns all offices from DB
func (o *Office) FindAllOffices(db *gorm.DB) (*[]Office, error) {
	var err error
	offices := []Office{}
	err = db.Debug().Model(&Office{}).Find(&offices).Error
	if err != nil {
		return &[]Office{}, err
	}
	return &offices, nil
}

// FindOfficeByID returns one office
func (o *Office) FindOfficeByID(db *gorm.DB, id int) (*Office, error) {
	var err error
	err = db.Debug().Model(&Office{}).Where("id = ?", id).Take(&o).Error
	if err != nil {
		return &Office{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Office{}, errors.New("Office not found")
	}
	return o, nil
}

// UpdateOffice updates office data
func (o *Office) UpdateOffice(db *gorm.DB) (*Office, error) {
	var err error
	err = db.Debug().Model(&Office{}).Where("id = ?", o.ID).Updates(Office{
		Country:    o.Country,
		City:       o.City,
		ZipCode:    o.ZipCode,
		Street:     o.Street,
		BuildingNr: o.BuildingNr,
		CompanyID:  o.CompanyID,
	}).Error
	if err != nil {
		return &Office{}, err
	}
	err = db.Debug().Model(&Office{}).Where("id = ?", o.ID).Take(&o).Error
	if err != nil {
		return &Office{}, err
	}
	return o, nil
}

// DeleteOffice removes office from DB
func (o *Office) DeleteOffice(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&Office{}).Where("id = ?", id).Take(&Office{}).Delete(&Office{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Office not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
