package usecase

import (
	"surunavi/pkg/domain"
	_const "surunavi/pkg/usecase/const"
	"surunavi/pkg/usecase/interfaces"
)

type LoginInteractor struct {
		LoginRepository interfaces.LoginRepository
}

func (interactor *LoginInteractor) Login(req domain.UserInfo) (_const.LoginResultType, domain.UserInfo, error) {

	userInfo, err := interactor.LoginRepository.GetUserInfo(req.Id)

	if err != nil {
		return _const.DbError, domain.UserInfo{}, err
	}

	if userInfo.Id == ""{
		return _const.NoUser,userInfo, nil
	} else if userInfo.Password != req.Password {
		return _const.MissPassword,userInfo, nil
	}
	return _const.Success,userInfo, nil
}
