package usecase

import (
	"surunavi/go/pkg/domain"
	_const "surunavi/go/pkg/domain/const"
	"surunavi/go/pkg/usecase/interfaces"
)

type LoginInteractor struct {
		LoginRepository interfaces.LoginRepository
}

func (interactor *LoginInteractor) Login(req domain.UserInfo) (_const.LoginResultType,domain.UserInfo) {

	userInfo := interactor.LoginRepository.GetUserInfo(req.Id)

	if userInfo.Id == ""{
		return _const.NoUser,userInfo
	} else if userInfo.Password != req.Password {
		return _const.MissPassword,userInfo
	}
	return _const.Succese,userInfo
}
