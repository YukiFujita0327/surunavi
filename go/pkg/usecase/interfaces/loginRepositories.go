package interfaces

import "surunavi/go/pkg/domain"

type LoginRepository interface {
	GetUserInfo(string) (domain.UserInfo, error)
}