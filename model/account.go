package model

// Account is struct of account
type Account struct {
	Base
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Level    int    `gorm:"column:level" json:"level"`
	State    int    `gorm:"column:state" json:"state"`
}

// TableName is return table's name
func (Account) TableName() string {
	return "chat_account"
}

// Add is add account
func (a *Account) Add() (err error) {
	err = db.Create(a).Error
	return
}
