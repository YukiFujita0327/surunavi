package interfaces

import "surunavi/pkg/domain"

type LoginRepository interface {
	GetUserInfo(string) (domain.UserInfo, error)
}