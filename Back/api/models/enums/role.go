package enums

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Role describes user role
type Role struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

// FindAllRoles method returns all roles from DB
func (r *Role) FindAllRoles(db *gorm.DB) (*[]Role, error) {
	var err error
	roles := []Role{}
	err = db.Debug().Model(&Role{}).Find(&roles).Error
	if err != nil {
		return &[]Role{}, err
	}
	return &roles, err
}

// FindRoleByID returns only one job experience that maches given ID
func (r *Role) FindRoleByID(db *gorm.DB, id int) (*Role, error) {
	var err error
	err = db.Debug().Model(&Role{}).Where("id = ?", id).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Role{}, errors.New("Role Not Found")
	}
	return r, err
}
