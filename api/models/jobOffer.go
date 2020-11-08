package models

import (
	"errors"
	"strings"
	"time"

	"github.com/abzibzi/jobOffers_API/api/models/enums"
	"github.com/jinzhu/gorm"
)

// JobOffer structure
type JobOffer struct {
	ID                   int                           `gorm:"primary_key;auto_increment" json:"id"`
	Name                 string                        `gorm:"size:255;not null" json:"name"`
	Description          string                        `gorm:"not null" json:"description"`
	SalaryMin            int                           `json:"salary_min"`
	SalaryMax            int                           `json:"salary_max"`
	ExperienceID         int                           `gorm:"not null" json:"experience_id"`
	Experience           enums.JobExperience           `gorm:"foreignKey:ExperienceID" json:"experience"`
	PublicationTime      time.Time                     `gorm:"default:CURRENT_TIMESTAMP" json:"publication_time"`
	CompanyID            int                           `gorm:"not null" json:"company_id"`
	Company              Company                       `gorm:"foreignKey:CompanyID" json:"company"`
	ContractTypeID       int                           `gorm:"not null" json:"contract_type_id"`
	ContractType         enums.CompanyType             `gorm:"foreignKey:ContractTypeID" json:"contract_type"`
	ExpectedTechnologies []enums.ExpectedTechnologyDTO `gorm:"many2many:job_expected_technologies" json:"expected_technologies"`
}

// Prepare func removes all white space before saving
func (j *JobOffer) Prepare() {
	j.ID = 0
	j.Name = strings.TrimSpace(j.Name)
	j.Description = strings.TrimSpace(j.Description)
	j.PublicationTime = time.Now()
	j.Company = Company{}
	j.ExpectedTechnologies = []enums.ExpectedTechnologyDTO{}
}

// Validate func checks if given data is valid
func (j *JobOffer) Validate() error {
	if j.Name == "" {
		return errors.New("Required Name")
	}
	if j.Description == "" {
		return errors.New("Required Description")
	}
	if j.SalaryMin < 0 {
		return errors.New("Required Minimal Salary")
	}
	if j.SalaryMax < 0 {
		return errors.New("Required Maximum Salary")
	}
	if j.SalaryMin > j.SalaryMax {
		return errors.New("Minimal salary is bigger than maximum salary")
	}
	if j.SalaryMax > 50000 || j.SalaryMin > 49999 {
		return errors.New("Salary is too big")
	}
	if j.ExperienceID < 0 {
		return errors.New("Required ExperienceID")
	}
	if j.CompanyID < 0 {
		return errors.New("Required CompanyID")
	}
	if j.ContractTypeID < 0 {
		return errors.New("Required ContractTypeID")
	}
	return nil
}

// SaveJobOffert func saves jobOffert to DB
func (j *JobOffer) SaveJobOffert(db *gorm.DB) (*JobOffer, error) {
	var err error
	err = db.Debug().Model(&JobOffer{}).Create(&j).Error
	if err != nil {
		return &JobOffer{}, err
	}
	if j.ID != 0 {
		err = db.Debug().Model(&enums.JobExperience{}).Where("id = ?", j.ExperienceID).Take(&j.Experience).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&Company{}).Where("id = ?", j.CompanyID).Take(&j.Company).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&enums.ContractType{}).Where("id = ?", j.ContractTypeID).Take(&j.ContractType).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&j).Related(&j.ExpectedTechnologies, "Technologies").Error
		if err != nil {
			return &JobOffer{}, err
		}
	}
	return j, nil
}

// FindAllJobOffers returns all job offers from DB
func (j *JobOffer) FindAllJobOffers(db *gorm.DB) (*[]JobOffer, error) {
	var err error
	jobs := []JobOffer{}
	err = db.Debug().Model(&Company{}).Find(&jobs).Error
	if err != nil {
		return &[]JobOffer{}, err
	}
	if len(jobs) > 0 {
		for i := range jobs {
			err = db.Debug().Model(&enums.JobExperience{}).Where("id = ?", jobs[i].ExperienceID).Take(&jobs[i].Experience).Error
			if err != nil {
				return &[]JobOffer{}, err
			}
			err = db.Debug().Model(&Company{}).Where("id = ?", jobs[i].CompanyID).Take(&jobs[i].Company).Error
			if err != nil {
				return &[]JobOffer{}, err
			}
			err = db.Debug().Model(&enums.ContractType{}).Where("id = ?", jobs[i].ContractTypeID).Take(&jobs[i].ContractType).Error
			if err != nil {
				return &[]JobOffer{}, err
			}
			technology := enums.JobExpectedTechnology{}
			technologies, err := technology.FindAllJobOfferTechnologies(db, jobs[i].ID)
			if err != nil {
				return &[]JobOffer{}, err
			}
			jobs[i].ExpectedTechnologies = technologies
		}
	}
	return &jobs, nil
}

// FindJobOffertByID func finds job offert by its ID in the BD
func (j *JobOffer) FindJobOffertByID(db *gorm.DB, id int) (*JobOffer, error) {
	var err error
	err = db.Debug().Model(&JobOffer{}).Where("id = ?", id).Take(&j).Error
	if err != nil {
		return &JobOffer{}, err
	}
	if j.ID != 0 {
		err = db.Debug().Model(&enums.JobExperience{}).Where("id = ?", j.ExperienceID).Take(&j.Experience).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&Company{}).Where("id = ?", j.CompanyID).Take(&j.Company).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&enums.ContractType{}).Where("id = ?", j.ContractTypeID).Take(&j.ContractType).Error
		if err != nil {
			return &JobOffer{}, err
		}
		technology := enums.JobExpectedTechnology{}
		technologies, err := technology.FindAllJobOfferTechnologies(db, j.ID)
		if err != nil {
			return &JobOffer{}, err
		}
		j.ExpectedTechnologies = technologies
	}
	return j, nil
}

// UpdateJobOffert saves new job offert data to DB
func (j *JobOffer) UpdateJobOffert(db *gorm.DB) (*JobOffer, error) {
	var err error
	err = db.Debug().Model(&JobOffer{}).Where("id = ?", j.ID).Updates(JobOffer{
		Name:           j.Name,
		Description:    j.Description,
		SalaryMin:      j.SalaryMin,
		SalaryMax:      j.SalaryMax,
		ExperienceID:   j.ExperienceID,
		CompanyID:      j.CompanyID,
		ContractTypeID: j.ContractTypeID,
	}).Error
	if err != nil {
		return &JobOffer{}, err
	}
	if j.ID != 0 {
		err = db.Debug().Model(&enums.JobExperience{}).Where("id = ?", j.ExperienceID).Take(&j.Experience).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&Company{}).Where("id = ?", j.CompanyID).Take(&j.Company).Error
		if err != nil {
			return &JobOffer{}, err
		}
		err = db.Debug().Model(&enums.ContractType{}).Where("id = ?", j.ContractTypeID).Take(&j.ContractType).Error
		if err != nil {
			return &JobOffer{}, err
		}
		technology := enums.JobExpectedTechnology{}
		technologies, err := technology.FindAllJobOfferTechnologies(db, j.ID)
		if err != nil {
			return &JobOffer{}, err
		}
		j.ExpectedTechnologies = technologies
	}
	return j, nil
}

// DeleteJobOffert removes job offert from DB
func (j *JobOffer) DeleteJobOffert(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&JobOffer{}).Where("id = ?", id).Take(&JobOffer{}).Delete(&JobOffer{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("JobOffer not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
