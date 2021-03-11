package interfaces

import (
	"surunavi/pkg/domain"
	_const "surunavi/pkg/usecase/const"
)

type LoginInteractor interface {
	Login(domain.UserInfo) (_const.LoginResultType, domain.UserInfo, error)
}
