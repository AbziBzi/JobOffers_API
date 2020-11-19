package enums

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// ContractType structure
type ContractType struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

// FindAllTypes method returns all types from DB
func (t *ContractType) FindAllTypes(db *gorm.DB) (*[]ContractType, error) {
	var err error
	types := []ContractType{}
	err = db.Debug().Model(&ContractType{}).Find(&types).Error
	if err != nil {
		return &[]ContractType{}, err
	}
	return &types, err
}

// FindTypeByID returns only one type that maches given ID
func (t *ContractType) FindTypeByID(db *gorm.DB, id int) (*ContractType, error) {
	var err error
	err = db.Debug().Model(&ContractType{}).Where("id = ?", id).Take(&t).Error
	if err != nil {
		return &ContractType{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &ContractType{}, errors.New("ContractType Not Found")
	}
	return t, err
}
