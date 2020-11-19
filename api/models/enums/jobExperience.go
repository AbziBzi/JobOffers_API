package enums

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// JobExperience structure
type JobExperience struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

// FindAllExperiences method returns all job experiences from DB
func (e *JobExperience) FindAllExperiences(db *gorm.DB) (*[]JobExperience, error) {
	var err error
	experiences := []JobExperience{}
	err = db.Debug().Model(&JobExperience{}).Find(&experiences).Error
	if err != nil {
		return &[]JobExperience{}, err
	}
	return &experiences, err
}

// FindExperienceByID returns only one job experience that maches given ID
func (e *JobExperience) FindExperienceByID(db *gorm.DB, id int) (*JobExperience, error) {
	var err error
	err = db.Debug().Model(&JobExperience{}).Where("id = ?", id).Take(&e).Error
	if err != nil {
		return &JobExperience{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &JobExperience{}, errors.New("JobExperience Not Found")
	}
	return e, err
}
