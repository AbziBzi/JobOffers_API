package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

// Company structure
type Company struct {
	ID           int      `gorm:"primary_key;auto_increment" json:"id"`
	Name         string   `gorm:"size:255;not null;unique" json:"name"`
	Size         int      `gorm:"not null" json:"size"`
	Industry     string   `gorm:"size:255; not null" json:"industry"`
	Headquarters string   `gorm:"size:255;not null" json:"headquarters"`
	SocialMedia  []string `json:"social_media"`
	TypeID       int      `gorm:"not null" json:"type"`
	Type         Type     `json:"company_type"`
}

// Prepare func removes all white space before saving
func (c *Company) Prepare() {
	c.ID = 0
	c.Name = strings.TrimSpace(c.Name)
	c.Industry = strings.TrimSpace(c.Industry)
	c.Headquarters = strings.TrimSpace(c.Headquarters)
	socialMedia := []string{}
	for _, mediaLink := range c.SocialMedia {
		socialMedia = append(socialMedia, strings.TrimSpace(mediaLink))
	}
	c.SocialMedia = socialMedia
}

// Validate func checks if given data is valid
func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("Required Name")
	}
	if c.Size < 1 {
		return errors.New("Required Size")
	}
	if c.Industry == "" {
		return errors.New("Required Industry")
	}
	if c.Headquarters == "" {
		return errors.New("Required Headquarters")
	}
	if c.TypeID < 1 {
		return errors.New("Required Type")
	}
	return nil
}

// SaveCompany fun saves company to DB
func (c *Company) SaveCompany(db *gorm.DB) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).Create(&c).Error
	if err != nil {
		return &Company{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&Type{}).Where("id = ?", c.TypeID).Take(&c.Type).Error
		if err != nil {
			return &Company{}, err
		}
	}
	return c, nil
}

// FindAllCompanies returns all companies from DB
func (c *Company) FindAllCompanies(db *gorm.DB) (*[]Company, error) {
	var err error
	companies := []Company{}
	err = db.Debug().Model(&Company{}).Find(&companies).Error
	if err != nil {
		return &[]Company{}, err
	}
	if len(companies) > 0 {
		for i := range companies {
			err := db.Debug().Model(&Type{}).Where("id = ?", companies[i].TypeID).Take(&companies[i].Type).Error
			if err != nil {
				return &[]Company{}, err
			}
		}
	}
	return &companies, nil
}

// FindCompanyByID func finds company by its ID in the BD
func (c *Company) FindCompanyByID(db *gorm.DB, id int) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).Where("id = ?", id).Take(&c).Error
	if err != nil {
		return &Company{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&Type{}).Where("id = ?", c.TypeID).Take(&c.Type).Error
		if err != nil {
			return &Company{}, err
		}
	}
	return c, nil
}

// UpdateCompany saves new company data to DB
func (c *Company) UpdateCompany(db *gorm.DB) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).Where("id = ?", c.ID).Updates(Company{
		Name:         c.Name,
		Size:         c.Size,
		Industry:     c.Industry,
		Headquarters: c.Headquarters,
		SocialMedia:  c.SocialMedia,
		TypeID:       c.TypeID,
	}).Error
	if err != nil {
		return &Company{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&Type{}).Where("id = ?", c.TypeID).Take(&c.Type).Error
		if err != nil {
			return &Company{}, err
		}
	}
	return c, nil
}

// DeleteCompany removes company object from DB
func (c *Company) DeleteCompany(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&Company{}).Where("id = ?", id).Take(&Company{}).Delete(&Company{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Company not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}