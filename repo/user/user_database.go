package user

import (
	"authentication/model"
	"gorm.io/gorm"
)

func NewUserDataBase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{db}
}

type UserDatabase struct{
	Db *gorm.DB
}

func (ud *UserDatabase) IsUser (user string, pass string) (bool){
	var userModel []model.BankUsrUsers
	result := ud.Db.Where("username=? and password=?",user,pass).Find(&userModel)
	if result.RowsAffected >0 {
		return true
	}else {
		return false
	}
}
