package dao

import (
	"fast-gin/model"
	"github.com/jinzhu/gorm"
	"log"
)

func GetAllUser(where map[string]interface{}, offset, limit int64) (users []*model.User, err error) {
	// 查询结构会直接填充到users里
	if limit > 0 && offset > 0 {
		err = db.Model(&model.User{}).Where(where).Find(&users).Offset(offset).Limit(limit).Error
	} else {
		err = db.Model(&model.User{}).Where(where).Find(&users).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetUser(id int64) (user *model.User, err error) {
	// 查询结构会直接填充到user里
	var u model.User
	err = db.Model(&model.User{}).Where("id = ?", id).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		log.Printf("GetUser  err(%v)", err)
		return
	}
	user = &u
	return
}

func CountUsers(where map[string]interface{}) (total int64, err error) {
	if err = db.Model(&model.User{}).Where(where).Count(&total).Error; err != nil {
		return
	}
	return
}

func AddUser(req model.User) (err error) {
	if err = db.Create(&req).Error; err != nil {
		return
	}
	return
}

func UpdateUser(id int64, data map[string]interface{}) (err error) {
	if err = db.Model(&model.User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return
	}
	return
}

func SaveUser(req model.User) (err error) {
	if err = db.Model(&model.User{}).Save(req).Error; err != nil {
		return
	}
	return
}

func DeleteUser(id int64) (err error) {
	if err = db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
