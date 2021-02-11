package interfaces

import "surunavi/go/pkg/domain"

type LoginInteractor interface {
	Login(domain.UserInfo) bool
}
