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
}

// Register is account register
func (a *Account) Register() error {
	var account model.Account
	account.Username = a.Username
	account.Password = a.Password
	account.Nickname = a.Nickname
	account.Salt = util.GenerateSalt(32)
	account.Level = 1
	account.State = 1
	return account.Add()
}
