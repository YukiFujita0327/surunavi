package usecase

import (
	"surunavi/go/pkg/domain"
	"surunavi/go/pkg/usecase/interfaces"
)

type LoginInteractor struct {
		LoginRepository interfaces.LoginRepository
}

func (interactor *LoginInteractor) Login(req domain.UserInfo) bool {
	interactor.LoginRepository.GetUserInfo(req.Id)
	return true
}
