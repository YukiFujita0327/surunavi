package interfaces

import "surunavi/go/pkg/domain"

type LoginInterracter interface{
	Login(domain.Login)(bool,error)
}

