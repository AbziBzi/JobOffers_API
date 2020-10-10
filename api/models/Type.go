package models

import (
	"errors"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// CompanyType struct that defines company type
type CompanyType struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

// Prepare func removes type input of any white space
func (t *CompanyType) Prepare() {
	t.ID = 0
	t.Name = strings.TrimSpace(t.Name)
}

// Validate method checks given data
func (t *CompanyType) Validate(action string) error {
	if t.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

// SaveType method saves type to DB
func (t *CompanyType) SaveType(db *gorm.DB) (*CompanyType, error) {
	var err error
	err = db.Debug().Create(&t).Error
	if err != nil {
		return &CompanyType{}, err
	}
	return t, nil
}

// FindAllTypes method returns all types from DB
func (t *CompanyType) FindAllTypes(db *gorm.DB) (*[]CompanyType, error) {
	var err error
	types := []CompanyType{}
	err = db.Debug().Model(&CompanyType{}).Find(&types).Error
	if err != nil {
		return &[]CompanyType{}, err
	}
	return &types, err
}

// FindTypeByID returns only one type that maches given ID
func (t *CompanyType) FindTypeByID(db *gorm.DB, id int) (*CompanyType, error) {
	var err error
	err = db.Debug().Model(&CompanyType{}).Where("id = ?", id).Take(&t).Error
	if err != nil {
		return &CompanyType{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CompanyType{}, errors.New("CompanyType Not Found")
	}
	return t, err
}

// UpdateType overwrite type's date
func (t *CompanyType) UpdateType(db *gorm.DB, id int) (*CompanyType, error) {
	var err error
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CompanyType{}).Where("id = ?", id).Take(&CompanyType{}).UpdateColumns(
		map[string]interface{}{
			"name": t.Name,
		},
	)
	if db.Error != nil {
		return &CompanyType{}, err
	}
	return t, nil
}

// DeleteType removes type from DB
func (t *CompanyType) DeleteType(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&CompanyType{}).Where("id = ?", id).Take(&CompanyType{}).Delete(&CompanyType{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
