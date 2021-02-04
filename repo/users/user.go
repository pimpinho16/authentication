package users

type BankUsrUsers struct{
	ID uint `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (BankUsrUsers) TableName() string {
	return "users.bank_usr_users"
}
