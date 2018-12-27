package app

import (
	"net/http"

	"chat/pkg/e"

	"github.com/astaxie/beego/validation"
)

// BindAndValid bind form and valid form data
func (g *Gin) BindAndValid(form interface{}) (int, int) {
	err := g.C.Bind(form)
	if nil != err {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if nil != err {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS
}
