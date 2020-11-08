package enums

import "github.com/jinzhu/gorm"

// JobExpectedTechnology struct
type JobExpectedTechnology struct {
	JobOfferID   int                      `json:"job_offer_id"`
	TechnologyID int                      `json:"technology_id"`
	Technology   Technology               `json:"technology"`
	ExperienceID int                      `json:"experience_id"`
	Experience   TechnologyExperienceType `json:"experience"`
}

// ExpectedTechnologyDTO for trying something new
type ExpectedTechnologyDTO struct {
	TechnologyName string `json:"tech_name"`
	Experience     string `json:"exp_name"`
}

// FindAllJobOfferTechnologies returns all Technologies that are related to given job offer
func (e *JobExpectedTechnology) FindAllJobOfferTechnologies(db *gorm.DB, jobOfferID int) ([]ExpectedTechnologyDTO, error) {
	var err error
	technologies := []JobExpectedTechnology{}
	returnTechnologies := []ExpectedTechnologyDTO{}
	err = db.Debug().Model(&JobExpectedTechnology{}).Where("joboffer_id = ?", jobOfferID).Find(&technologies).Error
	if err != nil {
		return []ExpectedTechnologyDTO{}, err
	}
	if len(technologies) > 0 {
		for i := range technologies {
			err = db.Debug().Model(&TechnologyExperienceType{}).Where("id = ?", technologies[i].ExperienceID).Take(&technologies[i].Experience).Error
			if err != nil {
				return []ExpectedTechnologyDTO{}, err
			}
			err = db.Debug().Model(&Technology{}).Where("id = ?", technologies[i].TechnologyID).Take(&technologies[i].Technology).Error
			if err != nil {
				return []ExpectedTechnologyDTO{}, err
			}
			returnTechnologies = append(returnTechnologies,
				ExpectedTechnologyDTO{
					TechnologyName: technologies[i].Technology.Name,
					Experience:     technologies[i].Experience.Name,
				})
		}
	}
	return returnTechnologies, nil
}
