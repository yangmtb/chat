package accountservice

import (
	"chat/model"
	"chat/pkg/util"
	"fmt"
)

// Param account param
type Param struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
}

// Account for form param
type Account struct {
	Params   Param
	Token    string
	haveInfo bool
	account  model.Account
}

// Signup is account register
func (a *Account) Signup() error {
	a.account.Username = a.Params.Username
	a.account.Nickname = a.Params.Nickname
	a.account.Email = a.Params.Email
	a.account.Phone = a.Params.Phone
	a.account.Salt = util.GenerateSalt(32)
	a.account.Password = util.GeneratePassword(a.Params.Password, a.account.Salt)
	a.account.Level = 1
	a.account.State = 1
	return a.account.Add()
}

// Auth password is ok
func (a *Account) Auth() bool {
	if !a.haveInfo {
		err := a.GetInfo()
		if nil != err {
			return false
		}
	}
	if util.GeneratePassword(a.Params.Password, a.account.Salt) == a.account.Password {
		fmt.Println("auth")
		return true
	}
	return false
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
