package model

import (
	"github.com/jinzhu/gorm"
)

// Account is struct of account
type Account struct {
	Base
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Salt     string `gorm:"column:salt" json:"salt"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`
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

// GetInfo is chekc username and password
func (a *Account) GetInfo() (err error) {
	err = db.Select("*").Where(*a).First(a).Error
	if nil != err && gorm.ErrRecordNotFound != err {
		return err
	}
	return nil
}
