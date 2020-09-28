package models

import (
	"errors"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// Type struct that defines company type
type Type struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

// Prepare func removes type input of any white space
func (t *Type) Prepare() {
	t.ID = 0
	t.Name = strings.TrimSpace(t.Name)
}

// Validate method checks given data
func (t *Type) Validate(action string) error {
	if t.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

// SaveType method saves type to DB
func (t *Type) SaveType(db *gorm.DB) (*Type, error) {
	var err error
	err = db.Debug().Create(&t).Error
	if err != nil {
		return &Type{}, err
	}
	return t, nil
}

// FindAllTypes method returns all types from DB
func (t *Type) FindAllTypes(db *gorm.DB) (*[]Type, error) {
	var err error
	types := []Type{}
	err = db.Debug().Model(&Type{}).Find(&types).Error
	if err != nil {
		return &[]Type{}, err
	}
	return &types, err
}

// FindTypeByID returns only one type that maches given ID
func (t *Type) FindTypeByID(db *gorm.DB, id int) (*Type, error) {
	var err error
	err = db.Debug().Model(&Type{}).Where("id = ?", id).Take(&t).Error
	if err != nil {
		return &Type{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Type{}, errors.New("Type Not Found")
	}
	return t, err
}

// UpdateType overwrite type's date
func (t *Type) UpdateType(db *gorm.DB, id int) (*Type, error) {
	var err error
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Type{}).Where("id = ?", id).Take(&Type{}).UpdateColumns(
		map[string]interface{}{
			"name": t.Name,
		},
	)
	if db.Error != nil {
		return &Type{}, err
	}
	return t, nil
}

// DeleteType removes type from DB
func (t *Type) DeleteType(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&Type{}).Where("id = ?", id).Take(&Type{}).Delete(&Type{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
