package gateways

import (
	"github.com/jinzhu/gorm"
	"surunavi/go/pkg/domain"
)

type (
	LoginRepository struct {
		Conn *gorm.DB
	}
)

func (repository *LoginRepository) GetUserInfo(userId string) domain.UserInfo {
	repository.Conn.Get(userId)
	return domain.UserInfo{Id: "", Password: ""}
}