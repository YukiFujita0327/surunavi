package gateways

import (
	"gorm.io/gorm"
	"surunavi/go/pkg/adapter/gateways/db"
	"surunavi/go/pkg/domain"
)

type (
	LoginRepository struct {
		Conn *gorm.DB
	}
)

func (repository *LoginRepository) GetUserInfo(userId string) domain.UserInfo {
	userInfo := &db.UserInfoModel{}
	repository.Conn.Where("Id = ?", userId).First(&userInfo)
	return domain.UserInfo{Id: userInfo.Id, Password: userInfo.Password, Name: userInfo.Name}
}