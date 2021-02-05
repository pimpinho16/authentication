package user

import (
	"authentication/model"
	"gorm.io/gorm"
	"log"
)

func NewUserDataBase(db *gorm.DB,logger log.Logger) *UserDatabase {
	return &UserDatabase{db,logger}
}

type UserDatabase struct{
	Db *gorm.DB
	logger log.Logger
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
