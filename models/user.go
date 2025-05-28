package model

import (
	"github.com/movie-rental-app/errors"
	"github.com/movie-rental-app/util"
)

type User struct {
	Base
	Name         string   `json:"name" gorm:"type:VARCHAR(100)"`
	Age          int      `json:"age" gorm:"type:INT"`
	Email        string   `json:"email" gorm:"type:VARCHAR(100)"`
	Password     string   `json:"password" gorm:"type:VARCHAR(100)"`
	RentedMovies []*Movie `json:"rentedMovies" gorm:"association_autocreate:false;association_autoupdate:false;"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Validate() error {

	if !util.IsValidEmailString(u.Email) {
		return errors.NewValidationError("email doesn't contain @")
	}

	if util.IsStringEmpty(u.Name) {
		return errors.NewValidationError("Name is empty.")
	}

	if u.Age < 0 || u.Age > 150 {
		return errors.NewValidationError("Age is less than 0 or greater than 150.")
	}
	return nil
}
