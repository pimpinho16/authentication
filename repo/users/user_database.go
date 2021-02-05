package users

import (
	"gorm.io/gorm"
)

func NewUserDataBase(db *gorm.DB) *UserDatabase {
	return &UserDatabase{db}
}

type UserDatabase struct{
	Db *gorm.DB
}

func (ud *UserDatabase) IsUser (user string, pass string) (BankUsrUsers,error){
	var userModel []BankUsrUsers
	error := ud.Db.Where("username=? and password=?",user,pass).Find(&userModel).Error
	return userModel[0],error
}
