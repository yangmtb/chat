package accountservice

import (
	"chat/model"
	"chat/pkg/util"
)

// Account for form param
type Account struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	Token    string
	haveInfo bool
	account  model.Account
}

// Register is account register
func (a *Account) Register() error {
	a.account.Username = a.Username
	a.account.Nickname = a.Nickname
	a.account.Email = a.Email
	a.account.Phone = a.Phone
	a.account.Salt = util.GenerateSalt(32)
	a.account.Password = util.GeneratePassword(a.Password, a.account.Salt)
	a.account.Level = 1
	a.account.State = 1
	return a.account.Add()
}

// Auth password is ok
func (a *Account) Auth() bool {
	if util.GeneratePassword(a.Password, a.account.Salt) == a.account.Password {
		return true
	} else {
		return false
	}
}

// GetInfo ...
func (a *Account) GetInfo() (err error) {
	err = a.account.GetInfo()
	if nil == err {
		a.haveInfo = true
	} else {
		a.haveInfo = false
	}
	return
}
