package dao

import (
	"fast-gin/model"
	"github.com/jinzhu/gorm"
)

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth model.Auth
	err := db.Select("id").Where(model.Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
