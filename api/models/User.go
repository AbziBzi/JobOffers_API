package models

import (
	"errors"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User model structure
type User struct {
	ID       int    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:100;not null" json:"name"`
	Surname  string `gorm:"size:100;not null" json:"surname"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
}

// Hash method is hashing user password
// it is made to protect user acc
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// BeforeSave method changes user password to hashed one
// This func should be performet before saving to database
// otherwise password will be saved not hashed
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare func removes user input of any white space
func (u *User) Prepare() {
	u.ID = 0
	u.Name = strings.TrimSpace(u.Name)
	u.Surname = strings.TrimSpace(u.Surname)
	u.Email = strings.TrimSpace(u.Email)
}

// Validate method checks given data
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Surname == "" {
			return errors.New("Required Surname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Surname == "" {
			return errors.New("Required Surname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// SaveUser method saves user to DB
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// FindAllUsers method returns all users from DB
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindUserByID returns only one user that maches given ID
func (u *User) FindUserByID(db *gorm.DB, id int) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

// UpdateUser overwrite user's date
func (u *User) UpdateUser(db *gorm.DB, id int) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password": u.Password,
			"name":     u.Name,
			"surname":  u.Surname,
			"email":    u.Email,
		},
	)
	if db.Error != nil {
		return &User{}, err
	}
	return u, nil
}

// DeleteUser removes user from DB
func (u *User) DeleteUser(db *gorm.DB, id int) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
