package interfaces

import (
	"surunavi/go/pkg/domain"
	_const "surunavi/go/pkg/usecase/const"
)

type LoginInteractor interface {
	Login(domain.UserInfo) (_const.LoginResultType,domain.UserInfo, error)
}
