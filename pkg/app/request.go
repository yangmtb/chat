package app

import (
	"chat/pkg/logging"

	"github.com/astaxie/beego/validation"
)

// MarkErrors log errors
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Debug(err.Key, err.Message)
	}
	return
}
