package dao

import (
	"github.com/coder9527-lg/lg-gin-vue/db"
	"github.com/coder9527-lg/lg-gin-vue/model"
)

func GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	// db := db.DB.First(&user)
	if err := db.DB.Debug().Where("phone=?", phone).Find(&user).Error; err != nil {
		return nil, err
	}
	// fmt.Println("db=", db)
	return user, nil
}
